# ![Image](/resources/render-logo-small.png) Render Redeploy Action

[![GitHub Super-Linter](https://github.com/qbaware/render-redeploy-action/actions/workflows/linter.yml/badge.svg)](https://github.com/super-linter/super-linter)
![CI](https://github.com/qbaware/render-redeploy-action/actions/workflows/ci.yml/badge.svg)
![CD](https://github.com/qbaware/render-redeploy-action/actions/workflows/cd.yml/badge.svg)

This is a simple GitHub Action that redeploys services hosted on
[Render](https://render.com/)'s cloud platform.

This action could be useful in CI/CD pipelines that release Docker images
and, as part of the pipeline, it's required to redeploy the services that
are using those images so that they can run the latest versions.

## How To Use

### Define The Inputs

First, make sure to define the necessary input for the Action. You'd
have to provide a Render API key (for authentication with their APIs)
and a Render service ID (used to identify the service you want to redeploy).

To do this, you can leverage GitHub
[secrets](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions)
and
[variables](https://docs.github.com/en/actions/learn-github-actions/variables).
In short, go to your repository's `Settings` tab and add the necessary
secrets and variables under `Secrets and variables` then `Actions`.

### Modify Your Action

Add the following `deploy` job in your Action.

``` yaml
deploy:
  needs: docker-images # Assuming this job builds the Docker images, we define a dependency on it.
  runs-on: ubuntu-latest
  steps:
    - name: Deploy to Render
      uses: qbaware/render-redeploy-action@v0
      with:
        render-api-key: ${{ secrets.RENDER_API_KEY }}
        render-service-id: ${{ env.RENDER_SERVICE_ID }}
```

### That's It 🎉
