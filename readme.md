
# Usage Examples for GitHub Actions Workflows

This README provides usage information for the various GitHub Actions workflow YAML examples.

## Docker Build and Push

The `docker-build.yml` workflow does the following:

- Builds a Docker image from the Dockerfile in the repo
- Tags the image with the GitHub repository name 
- Pushes the image to GitHub Container Registry

To use this:

- Set the `DOCKER_REGISTRY` env variable to your Docker registry URL
- Set the `IMAGE_NAME` env variable to your desired image name
- Update the Dockerfile path if needed
- Add Docker login credentials to your repo secrets

## Deploy via SSH

The `deploy-ssh.yml` workflow does the following:

- Installs the SSH key from a secret into the workflow
-SSHes to a remote server to run deployment scripts

To use this:

- Add your private SSH key as a secret called `SSH_PRIVATE_KEY`
- Set the `USERNAME`, `HOST`, `DEPLOY_PATH`, and `PORT` env variables
- Update the deploy scripts to match your app

## Commit and Workflow Notifications

The `notifications.yml` workflow does the following:

- Sends a Telegram notification on new commits
- Sends a Telegram notification when the workflow completes

To use this:

- Add your Telegram bot token as the secret `TELEGRAM_TOKEN` 
- Add your Telegram chat ID as the secret `TELEGRAM_TO`
- Customize the notification messages as needed

Let me know if any part of the usage needs more explanation!