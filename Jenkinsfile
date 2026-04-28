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
                sh 'docker run --rm -v "$WORKSPACE:/app" -w /app golang:1.25-alpine go test -v ./...'
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
