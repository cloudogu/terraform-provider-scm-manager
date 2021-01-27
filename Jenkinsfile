#!groovy

node('docker') {

    branch = "${env.BRANCH_NAME}"

    stage('Checkout') {
        checkout scm
    }

    docker.image('golang:1.14.13').inside("-e GOCACHE=/tmp") {

        stage('Build') {
            make 'clean compile checksum'
            archiveArtifacts 'target/*'
        }

        stage('Unit Test') {
            make 'unit-test'
            junit allowEmptyResults: true, testResults: 'target/unit-tests/*-tests.xml'
        }

        stage('Static Analysis') {
            make 'static-analysis'
        }
    }

    stage('SonarQube') {
        def scannerHome = tool name: 'sonar-scanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
        withSonarQubeEnv {
            if (branch == "main") {
                echo "This branch has been detected as the main branch."
                sh "${scannerHome}/bin/sonar-scanner"
            } else if (branch == "develop") {
                echo "This branch has been detected as the develop branch."
                sh "${scannerHome}/bin/sonar-scanner -Dsonar.branch.name=${env.BRANCH_NAME} -Dsonar.branch.target=master"
            } else if (env.CHANGE_TARGET) {
                echo "This branch has been detected as a pull request."
                sh "${scannerHome}/bin/sonar-scanner -Dsonar.branch.name=${env.CHANGE_BRANCH}-PR${env.CHANGE_ID} -Dsonar.branch.target=${env.CHANGE_TARGET}"
            } else if (branch.startsWith("feature/")) {
                echo "This branch has been detected as a feature branch."
                sh "${scannerHome}/bin/sonar-scanner -Dsonar.branch.name=${env.BRANCH_NAME} -Dsonar.branch.target=develop"
            }
        }
    }
}

void make(String goal) {
    sh "make ${goal}"
}