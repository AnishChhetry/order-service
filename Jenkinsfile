pipeline {
    agent any

    environment {
        APP_NAME = 'order-service'
        DOCKER_IMAGE = "${APP_NAME}:v1"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test') {
            steps {
                // Use docker build instead of docker run -v to avoid Docker-in-Docker volume mount issues
                sh '''
                    docker build -f - . <<'EOF'
FROM golang:1.25-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test -v ./...
EOF
                '''
            }
        }

        stage('Docker Build') {
            steps {
                sh "docker build -t ${DOCKER_IMAGE} ."
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh "kubectl apply -f k8s.yaml"
            }
        }
    }
}
