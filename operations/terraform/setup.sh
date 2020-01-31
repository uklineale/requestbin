curl -o terraform.zip https://releases.hashicorp.com/terraform/0.12.19/terraform_0.12.19_linux_amd64.zip
unzip terraform.zip
sudo mv terraform /usr/bin/
sudo yum install -y jq