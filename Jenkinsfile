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
            go mod init  
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
  }
}
