package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsApiGatewayDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsApiGatewayDeploymentCreate,
		Read:   resourceAwsApiGatewayDeploymentRead,
		Update: resourceAwsApiGatewayDeploymentUpdate,
		Delete: resourceAwsApiGatewayDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"rest_api_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"stage_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"stage_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"variables": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"created_date": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"invoke_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"execution_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"xray_tracing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsApiGatewayDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).apigateway
	// Create the gateway
	log.Printf("[DEBUG] Creating API Gateway Deployment")

	input := &apigateway.CreateDeploymentInput{
		RestApiId: aws.String(d.Get("rest_api_id").(string)),
		StageName: aws.String(d.Get("stage_name").(string)),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}
	if v, ok := d.GetOk("stage_description"); ok {
		input.StageDescription = aws.String(v.(string))
	}
	if v, ok := d.GetOk("xray_tracing_enabled"); ok {
		input.TracingEnabled = aws.Bool(v.(bool))
	}
	if v, ok := d.GetOk("variables"); ok {
		variables := make(map[string]string)
		for k, v := range v.(map[string]interface{}) {
			variables[k] = v.(string)
		}
		input.Variables = aws.StringMap(variables)
	}

	deployment, err := conn.CreateDeployment(input)
	if err != nil {
		return fmt.Errorf("Error creating API Gateway Deployment: %s", err)
	}

	d.SetId(*deployment.Id)
	log.Printf("[DEBUG] API Gateway Deployment ID: %s", d.Id())

	return resourceAwsApiGatewayDeploymentRead(d, meta)
}

func resourceAwsApiGatewayDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).apigateway

	log.Printf("[DEBUG] Reading API Gateway Deployment %s", d.Id())
	restApiId := d.Get("rest_api_id").(string)
	out, err := conn.GetDeployment(&apigateway.GetDeploymentInput{
		RestApiId:    aws.String(restApiId),
		DeploymentId: aws.String(d.Id()),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NotFoundException" {
			log.Printf("[WARN] API Gateway Deployment (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}
	log.Printf("[DEBUG] Received API Gateway Deployment: %s", out)
	d.Set("description", out.Description)

	region := meta.(*AWSClient).region
	stageName := d.Get("stage_name").(string)

	d.Set("invoke_url", buildApiGatewayInvokeURL(restApiId, region, stageName))

	executionArn := arn.ARN{
		Partition: meta.(*AWSClient).partition,
		Service:   "execute-api",
		Region:    meta.(*AWSClient).region,
		AccountID: meta.(*AWSClient).accountid,
		Resource:  fmt.Sprintf("%s/%s", restApiId, stageName),
	}.String()
	d.Set("execution_arn", executionArn)

	if err := d.Set("created_date", out.CreatedDate.Format(time.RFC3339)); err != nil {
		log.Printf("[DEBUG] Error setting created_date: %s", err)
	}

	input := &apigateway.GetStageInput{
		RestApiId: aws.String(d.Get("rest_api_id").(string)),
		StageName: aws.String(d.Get("stage_name").(string)),
	}
	stage, err := conn.GetStage(input)
	if err != nil {
		return err
	}
	d.Set("xray_tracing_enabled", stage.TracingEnabled)
	d.Set("stage_description", stage.Description)
	d.Set("stage_name", stage.StageName)
	if err := d.Set("variables", aws.StringValueMap(stage.Variables)); err != nil {
		return fmt.Errorf("error setting stage variables: %s", err)
	}

	return nil
}

func resourceAwsApiGatewayDeploymentUpdateOperations(d *schema.ResourceData) []*apigateway.PatchOperation {
	operations := make([]*apigateway.PatchOperation, 0)

	if d.HasChange("description") {
		operations = append(operations, &apigateway.PatchOperation{
			Op:    aws.String("replace"),
			Path:  aws.String("/description"),
			Value: aws.String(d.Get("description").(string)),
		})
	}

	return operations
}

func resourceAwsApiGatewayDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).apigateway

	log.Printf("[DEBUG] Updating API Gateway API Key: %s", d.Id())

	_, err := conn.UpdateDeployment(&apigateway.UpdateDeploymentInput{
		DeploymentId:    aws.String(d.Id()),
		RestApiId:       aws.String(d.Get("rest_api_id").(string)),
		PatchOperations: resourceAwsApiGatewayDeploymentUpdateOperations(d),
	})
	if err != nil {
		return err
	}

	return resourceAwsApiGatewayDeploymentRead(d, meta)
}

func resourceAwsApiGatewayDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).apigateway
	log.Printf("[DEBUG] Deleting API Gateway Deployment: %s", d.Id())

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		log.Printf("[DEBUG] schema is %#v", d)
		if _, err := conn.DeleteStage(&apigateway.DeleteStageInput{
			StageName: aws.String(d.Get("stage_name").(string)),
			RestApiId: aws.String(d.Get("rest_api_id").(string)),
		}); err == nil {
			return nil
		}

		_, err := conn.DeleteDeployment(&apigateway.DeleteDeploymentInput{
			DeploymentId: aws.String(d.Id()),
			RestApiId:    aws.String(d.Get("rest_api_id").(string)),
		})
		if err == nil {
			return nil
		}

		apigatewayErr, ok := err.(awserr.Error)
		if apigatewayErr.Code() == "NotFoundException" {
			return nil
		}

		if !ok {
			return resource.NonRetryableError(err)
		}

		return resource.NonRetryableError(err)
	})
}
