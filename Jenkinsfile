@Library("shared-library") _
pipeline {
    agent any
    stages {
        stage('Hello') {
            steps {
                helloWorldV1(name:"Name", dayOfWeek:"Day")
            }
        }
        stage('Run') {
            dir('Workspace01') {
                steps {
                    sh('cd kubernetes')
                    sh('cd jobs')
                    k8sDeployLocal('Job5')
                }
            }
        }
    }
}
