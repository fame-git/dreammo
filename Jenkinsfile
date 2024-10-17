pipeline {
  agent any

  environment {
    GOPATH = "/workspace/go"  // Set GOPATH to the same writable directory
  }

  stages {

    stage('Verify Agent'){
      agent {
        docker {
          image 'go-agent'
          reuseNode true
        }
      }

      steps {
        sh '''
          go version
          echo "------fin check go--------------"

          kubectl version --client
          echo "------fin kubectl verify--------"
        '''
      }
    }

    stage('Init build'){
      agent {
        docker {
          image 'go-agent'
          reuseNode true
        }
      }

      steps {
        sh '''
            go mod tidy
            go mod download  
        '''

      }
    }

    stage('Unit Test'){
      agent {
        docker {
          image 'go-agent'
          reuseNode true
        }
      }

      steps {
        sh '''
            go test        
        '''

      }
    }

    stage('Build GO'){
      agent {
        docker {
          image 'go-agent'
          reuseNode true
        }
      }

      steps {
        sh '''
            go build -o myapp .     
        '''

      }
    }

    stage('Docker build'){
      agent any

      environment {
        DOCKER_HUB_REPO = "famedocker/go-back"
        DOCKER_TAG = "${BUILD_NUMBER}"  // Dynamic tag based on branch name and build number
      }

      steps {
        sh 'docker build -t ${DOCKER_HUB_REPO}:${DOCKER_TAG} .'
      }
    }
  }
}
