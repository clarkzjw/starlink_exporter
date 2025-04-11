sudo apt-get install curl -y

curl -fsSL https://get.docker.com | bash
sudo systemctl enable docker
sudo systemctl restart docker

sudo docker compose pull
sudo docker compose up -d
