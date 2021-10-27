#!groovy
@Library(['github.com/cloudogu/zalenium-build-lib@v2.1.0'])
import com.cloudogu.ces.zaleniumbuildlib.*

node('docker') {

    branch = "${env.BRANCH_NAME}"

    stage('Checkout') {
        checkout scm
    }

    def scmImage = docker.image('scmmanager/scm-manager:2.12.0')
    def scmContainerName = "${JOB_BASE_NAME}-${BUILD_NUMBER}".replaceAll("\\/|%2[fF]", "-")
    withDockerNetwork { buildnetwork ->
        scmImage.withRun("--network ${buildnetwork} --name ${scmContainerName}") {

            docker.image('golang:1.14.13').inside("--network ${buildnetwork} -e HOME=/tmp") {

                stage('Build') {
                    make 'clean package checksum'
                    archiveArtifacts 'target/*'
                }

                stage('Unit Test') {
                    make 'unit-test'
                    junit allowEmptyResults: true, testResults: 'target/unit-tests/*-tests.xml'
                }

                stage('Static Analysis') {
                    make 'static-analysis'
                }

                stage('Acceptance Tests') {
                    sh "SCM_URL=http://${scmContainerName}:8080/scm make testacc"
                    archiveArtifacts 'target/acceptance-tests/*.out'
                }

            }
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
        timeout(time: 2, unit: 'MINUTES') { // Needed when there is no webhook for example
            def qGate = waitForQualityGate()
            if (qGate.status != 'OK') {
                unstable("Pipeline unstable due to SonarQube quality gate failure")
            }
        }
    }

    if (gitflow.isReleaseBranch()) {
        String releaseVersion = git.getSimpleBranchName();

        stage('Finish Release') {
            gitflow.finishRelease(releaseVersion)
        }

        stage('Add Github-Release') {
            github.createReleaseWithChangelog(releaseVersion, changelog)
        }
    }
}

void make(String goal) {
    sh "make ${goal}"
}