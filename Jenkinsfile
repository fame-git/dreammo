pipeline {
  agent any

  environment {
    DOCKER_HUB_REPO = "famedocker/go-back"
    DOCKER_TAG = "${BUILD_NUMBER}"  // Dynamic tag based on branch name and build number
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

      steps {
        sh 'docker build -t ${DOCKER_HUB_REPO}:${DOCKER_TAG} .'
      }
    }

    stage('Docker Publish'){
      agent any
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
          sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
          sh 'docker push ${DOCKER_HUB_REPO}:${DOCKER_TAG}'
        }
      }
    }
  }
}
