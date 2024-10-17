pipeline {
  agent any

  stages {

    stage('Build'){
      agent {
        docker {
          image 'golang:1.22.5-alpine'
          reuseNode true
        }
      }

      steps {
        sh '''
          go version

          # Set GOCACHE to a directory with appropriate permissions
          export GOCACHE=/tmp/.cache
          mkdir -p /tmp/.cache

          
        '''
      }
    }
  }
}
