################################################################################
# Create a VPC for isolating our traffic                                       #
################################################################################
resource "digitalocean_vpc" "web"{
    # The human friendly name of our VPC.
    name = "tourofpython-vpc"

    # The region to deploy our VPC to.
    region = var.region

    # The private ip range within our VPC
    ip_range = "192.168.31.0/24"
}
