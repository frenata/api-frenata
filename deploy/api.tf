provider "aws" {
  region = "us-east-2"
}

resource "aws_eip" "ip" {
  instance = "${aws_instance.api.id}"
}

resource "aws_security_group" "instance" {
  name = "api-instance"

  ingress {
    from_port = 80
    to_port   = 80
    protocol  =  "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "http"
  }

  ingress {
    from_port = 3000
    to_port   = 3000
    protocol  =  "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "http"
  }

  ingress {
    from_port = 22
    to_port   = 22
    protocol  =  "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "ssh"
  }

  egress {
    from_port = 0
    to_port   = 65535
    protocol  =  "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "api" {
  ami = "${lookup(var.amis, var.region)}"
  instance_type = "t2.micro"
  user_data = "${file("setup.sh")}"
  vpc_security_group_ids = ["${aws_security_group.instance.id}"]
  key_name = "${aws_key_pair.generated_key.key_name}"
}
