
# LGTM Bot

A GitHub Action that automatically adds a random LGTM (Looks Good To Me) image to an existing comment when it contains the text "lgtm" (case-insensitive). This bot uses images from the [LGTMeow API](https://lgtmeow.com/api/lgtm-images).

## Features

- Adds an LGTM image to comments that include "lgtm" (case-insensitive).
- Avoids duplicate LGTM images in the same comment.
- Works on both issues and pull requests.

## Requirements

- **GitHub Token**: Required to authenticate and modify comments on GitHub.
- **Go**: This action is implemented in Go, so ensure you have it set up in your environment if modifying the code locally.

## Usage

1. Create a workflow file in your repository (e.g., `.github/workflows/lgtm-workflow.yml`).
2. Copy the example below into your workflow file.

### Example Workflow

```yaml
name: LGTM Bot

on:
  issue_comment:
    types: [created]
  pull_request_review:
    types: [submitted]

jobs:
  lgtm:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v2

      - name: Check if comment contains 'lgtm' (case-insensitive) and no existing LGTM image
        id: check_comment
        run: |
          echo "Comment body: ${{ github.event.comment.body }}"
          COMMENT_BODY="${{ github.event.comment.body }}"
          COMMENT_BODY_LOWER=$(echo "$COMMENT_BODY" | tr '[:upper:]' '[:lower:]')
          if [[ "$COMMENT_BODY_LOWER" != *"lgtm"* || "$COMMENT_BODY" == *"![LGTM]"* ]]; then
            echo "skip=true" >> $GITHUB_ENV
          else
            echo "skip=false" >> $GITHUB_ENV
          fi

      - name: Run LGTM Bot Action
        if: env.skip == 'false'
        uses: db-r-hashimoto/lgtm_print@v0.0.1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          owner: ${{ github.repository_owner }}
          repo: ${{ github.event.repository.name }}
          comment_body: ${{ github.event.comment.body }}
          comment_id: ${{ github.event.comment.id }}
```

### Inputs

| Name            | Description                                | Required | Default                      |
|-----------------|--------------------------------------------|----------|------------------------------|
| `github-token`  | GitHub token for authentication            | Yes      | `${{ secrets.GITHUB_TOKEN }}`|
| `owner`         | Repository owner                           | Yes      |                              |
| `repo`          | Repository name                            | Yes      |                              |
| `comment_body`  | The content of the comment                 | Yes      |                              |
| `comment_id`    | The ID of the comment                      | Yes      |                              |

### Environment Variables

| Variable       | Description                                                                 |
|----------------|-----------------------------------------------------------------------------|
| `GITHUB_TOKEN` | Required to authenticate with the GitHub API for modifying comments.        |

### Setup Instructions

1. Go to **Settings** > **Secrets and variables** > **Actions** in your repository.
2. Add a new secret named `GITHUB_TOKEN` with your GitHub Personal Access Token (or use `${{ secrets.GITHUB_TOKEN }}` for default GitHub token).
3. Copy the workflow example above to your `.github/workflows` directory to enable the bot.

### License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

### Acknowledgements

- LGTM images are sourced from [LGTMeow API](https://lgtmeow.com/api/lgtm-images).
