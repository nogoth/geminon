# Deployment Guide: Geminon LCARS Dashboard

This project is designed to be hosted on **GitHub Pages** using **GitHub Actions**.

## Prerequisites
1. Push this project to a GitHub repository.
2. Ensure `vite.config.js` has `base: './'` or `base: '/<REPO_NAME>/'`.

## Step-by-Step Deployment

### 1. Enable GitHub Pages
- Go to your repository settings on GitHub.
- Navigate to **Pages** in the sidebar.
- Under **Build and deployment > Source**, select **GitHub Actions**.

### 2. Configure GitHub Actions
- The project includes a `.github/workflows/deploy.yml` file (created in Phase 5).
- Every time you push to the `main` (or `github-vue`) branch, the action will:
  - Install dependencies.
  - Build the project.
  - Deploy the static assets to the `gh-pages` branch.

### 3. Verify Deployment
- Once the Action completes, your dashboard will be live at `https://<USERNAME>.github.io/<REPO_NAME>/`.

## Technical Notes
- **Vite Base Path:** If you are hosting on a custom domain, set `base: '/'`. If hosting on `<username>.github.io/<repo>/`, use the repository name.
- **Node Version:** The workflow uses Node 20.x.
