package aws

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Sagemaker Algorithm BlazingText
	sageMakerAlgorithmBlazingText = "blazingtext"
	// Sagemaker Algorithm DeepAR Forecasting
	sageMakerAlgorithmDeepARForecasting = "forecasting-deepar"
	// Sagemaker Algorithm Factorization Machines
	sageMakerAlgorithmFactorizationMachines = "factorization-machines"
	// Sagemaker Algorithm Image Classification
	sageMakerAlgorithmImageClassification = "image-classification"
	// Sagemaker Algorithm IP Insights
	sageMakerAlgorithmIPInsights = "ipinsights"
	// Sagemaker Algorithm k-means
	sageMakerAlgorithmKMeans = "kmeans"
	// Sagemaker Algorithm k-nearest-neighbor
	sageMakerAlgorithmKNearestNeighbor = "knn"
	// Sagemaker Algorithm Latent Dirichlet Allocation
	sageMakerAlgorithmLDA = "lda"
	// Sagemaker Algorithm Linear Learner
	sageMakerAlgorithmLinearLearner = "linear-learner"
	// Sagemaker Algorithm Neural Topic Model
	sageMakerAlgorithmNeuralTopicModel = "ntm"
	// Sagemaker Algorithm Object2Vec
	sageMakerAlgorithmObject2Vec = "object2vec"
	// Sagemaker Algorithm Object Detection
	sageMakerAlgorithmObjectDetection = "object-detection"
	// Sagemaker Algorithm PCA
	sageMakerAlgorithmPCA = "pca"
	// Sagemaker Algorithm Random Cut Forest
	sageMakerAlgorithmRandomCutForest = "randomcutforest"
	// Sagemaker Algorithm Semantic Segmentation
	sageMakerAlgorithmSemanticSegmentation = "semantic-segmentation"
	// Sagemaker Algorithm Seq2Seq
	sageMakerAlgorithmSeq2Seq = "seq2seq"
	// Sagemaker Algorithm XGBoost
	sageMakerAlgorithmXGBoost = "sagemaker-xgboost"
)

var sageMakerPrebuiltImageIDByRegion_Blazing = map[string]string{
	"ap-east-1":      "286214385809",
	"ap-northeast-1": "501404015308",
	"ap-northeast-2": "306986355934",
	"ap-south-1":     "991648021394",
	"ap-southeast-1": "475088953585",
	"ap-southeast-2": "544295431143",
	"ca-central-1":   "469771592824",
	"cn-north-1":     "390948362332",
	"cn-northwest-1": "387376663083",
	"eu-central-1":   "813361260812",
	"eu-north-1":     "669576153137",
	"eu-west-1":      "685385470294",
	"eu-west-2":      "644912444149",
	"eu-west-3":      "749696950732",
	"me-south-1":     "249704162688",
	"sa-east-1":      "855470959533",
	"us-east-1":      "811284229777",
	"us-east-2":      "825641698319",
	"us-gov-west-1":  "226302683700",
	"us-west-1":      "632365934929",
	"us-west-2":      "433757028032",
}

var sageMakerPrebuiltImageIDByRegion_DeepAR = map[string]string{
	"ap-east-1":      "286214385809",
	"ap-northeast-1": "633353088612",
	"ap-northeast-2": "204372634319",
	"ap-south-1":     "991648021394",
	"ap-southeast-1": "475088953585",
	"ap-southeast-2": "514117268639",
	"ca-central-1":   "469771592824",
	"cn-north-1":     "390948362332",
	"cn-northwest-1": "387376663083",
	"eu-central-1":   "495149712605",
	"eu-north-1":     "669576153137",
	"eu-west-1":      "224300973850",
	"eu-west-2":      "644912444149",
	"eu-west-3":      "749696950732",
	"me-south-1":     "249704162688",
	"sa-east-1":      "855470959533",
	"us-east-1":      "522234722520",
	"us-east-2":      "566113047672",
	"us-gov-west-1":  "226302683700",
	"us-west-1":      "632365934929",
	"us-west-2":      "156387875391",
}

var sageMakerPrebuiltImageIDByRegion_FactorMachines = map[string]string{
	"ap-east-1":      "286214385809",
	"ap-northeast-1": "351501993468",
	"ap-northeast-2": "835164637446",
	"ap-south-1":     "991648021394",
	"ap-southeast-1": "475088953585",
	"ap-southeast-2": "712309505854",
	"ca-central-1":   "469771592824",
	"cn-north-1":     "390948362332",
	"cn-northwest-1": "387376663083",
	"eu-central-1":   "664544806723",
	"eu-north-1":     "669576153137",
	"eu-west-1":      "438346466558",
	"eu-west-2":      "644912444149",
	"eu-west-3":      "749696950732",
	"me-south-1":     "249704162688",
	"sa-east-1":      "855470959533",
	"us-east-1":      "382416733822",
	"us-east-2":      "404615174143",
	"us-gov-west-1":  "226302683700",
	"us-west-1":      "632365934929",
	"us-west-2":      "174872318107",
}

var sageMakerPrebuiltImageIDByRegion_LDA = map[string]string{
	"ap-northeast-1": "258307448986",
	"ap-northeast-2": "293181348795",
	"ap-south-1":     "991648021394",
	"ap-southeast-1": "475088953585",
	"ap-southeast-2": "297031611018",
	"ca-central-1":   "469771592824",
	"eu-central-1":   "353608530281",
	"eu-west-1":      "999678624901",
	"eu-west-2":      "644912444149",
	"us-east-1":      "766337827248",
	"us-east-2":      "999911452149",
	"us-gov-west-1":  "226302683700",
	"us-west-1":      "632365934929",
	"us-west-2":      "266724342769",
}

var sageMakerPrebuiltImageIDByRegion_XGBoost = map[string]string{
	"ap-east-1":      "651117190479",
	"ap-northeast-1": "354813040037",
	"ap-northeast-2": "366743142698",
	"ap-south-1":     "720646828776",
	"ap-southeast-1": "121021644041",
	"ap-southeast-2": "783357654285",
	"ca-central-1":   "341280168497",
	"cn-north-1":     "450853457545",
	"cn-northwest-1": "451049120500",
	"eu-central-1":   "492215442770",
	"eu-north-1":     "662702820516",
	"eu-west-1":      "141502667606",
	"eu-west-2":      "764974769150",
	"eu-west-3":      "659782779980",
	"me-south-1":     "801668240914",
	"sa-east-1":      "737474898029",
	"us-east-1":      "683313688378",
	"us-east-2":      "257758044811",
	"us-gov-west-1":  "414596584902",
	"us-west-1":      "746614075791",
	"us-west-2":      "246618743249",
}

func dataSourceAwsSagemakerPrebuiltImagePath() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsRdsOrderableDbInstanceRead,
		Schema: map[string]*schema.Schema{
			"algorithm_name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					sageMakerAlgorithmBlazingText,
					sageMakerAlgorithmDeepARForecasting,
					sageMakerAlgorithmFactorizationMachines,
					sageMakerAlgorithmImageClassification,
					sageMakerAlgorithmIPInsights,
					sageMakerAlgorithmKMeans,
					sageMakerAlgorithmKNearestNeighbor,
					sageMakerAlgorithmLDA,
					sageMakerAlgorithmLinearLearner,
					sageMakerAlgorithmNeuralTopicModel,
					sageMakerAlgorithmObject2Vec,
					sageMakerAlgorithmObjectDetection,
					sageMakerAlgorithmPCA,
					sageMakerAlgorithmRandomCutForest,
					sageMakerAlgorithmSemanticSegmentation,
					sageMakerAlgorithmSeq2Seq,
					sageMakerAlgorithmXGBoost,
				}, false),
			},

			"dns_suffix": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"registry_path": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},
		},
	}
}

func dataSourceAwsSagemakerPrebuiltImagePathRead(d *schema.ResourceData, meta interface{}) error {
	region := meta.(*AWSClient).region
	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	}

	suffix := meta.(*AWSClient).dnsSuffix
	if v, ok := d.GetOk("dns_suffix"); ok {
		suffix = v.(string)
	}

	algo := d.Get("algorithm_name").(string)
	tag := d.Get("tag").(string)

	if algo == sageMakerAlgorithmBlazingText || algo == sageMakerAlgorithmImageClassification || algo == sageMakerAlgorithmObjectDetection || algo == sageMakerAlgorithmSemanticSegmentation || algo == sageMakerAlgorithmSeq2Seq {
		if id, ok := sageMakerPrebuiltImageIDByRegion_Blazing[region]; ok {
			d.SetId(id)
			d.Set("registry_path", dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algo, tag))
			return nil
		}
		return fmt.Errorf("no image path available for region (%s) and algorithm (%s)", region, algo)
	}

	if algo == sageMakerAlgorithmDeepARForecasting {
		if id, ok := sageMakerPrebuiltImageIDByRegion_DeepAR[region]; ok {
			d.SetId(id)
			d.Set("registry_path", dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algo, tag))
			return nil
		}
		return fmt.Errorf("no image path available for region (%s) and algorithm (%s)", region, algo)
	}

	if algo == sageMakerAlgorithmLDA {
		if id, ok := sageMakerPrebuiltImageIDByRegion_LDA[region]; ok {
			d.SetId(id)
			d.Set("registry_path", dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algo, tag))
			return nil
		}
		return fmt.Errorf("no image path available for region (%s) and algorithm (%s)", region, algo)
	}

	if algo == sageMakerAlgorithmXGBoost {
		if id, ok := sageMakerPrebuiltImageIDByRegion_XGBoost[region]; ok {
			d.SetId(id)
			d.Set("registry_path", dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algo, tag))
			return nil
		}
		return fmt.Errorf("no image path available for region (%s) and algorithm (%s)", region, algo)
	}

	if id, ok := sageMakerPrebuiltImageIDByRegion_FactorMachines[region]; ok {
		d.SetId(id)
		d.Set("registry_path", dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algo, tag))
		return nil
	}
	return fmt.Errorf("no image path available for region (%s) and algorithm (%s)", region, algo)
}

func dataSourceAwsSagemakerPrebuiltImageCreatePath(id, region, suffix, algorithm, tag string) string {
	return fmt.Sprintf("%s.dkr.ecr.%s.%s/%s:%s", id, region, suffix, algorithm, tag)
}
