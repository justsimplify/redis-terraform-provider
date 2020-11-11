cleanup:
	rm -rf .terraform terraform.*

init:
	terraform init

plan:
	terraform plan -var-file=variables.tfvars

apply:
	terraform apply -var-file=variables.tfvars -auto-approve

setup: cleanup init plan apply

read_key: cleanup init apply

refresh:
	terraform refresh -var-file=variables.tfvars