pipeline {
  agent any

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
            # Set GOCACHE to a directory with appropriate permissions
            export GOCACHE=/tmp/.cache
            mkdir -p /tmp/.cache

            go mod init
        '''

      }
    }
  }
}
