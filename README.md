# Clone Organisation Repoistories

Clones or updates

## Usage

- Set a GITHUB_TOKEN environment variable to one which can pull from all the repositories you need
- Set the GITHUB_ORG environment variable to the org to clone from
- Run clone-org-repos

```bash
$ GITHUB_ORG="sous-chefs" ./clone-org-repos

Cloning repository: rvm
SSH CLone URL: git@github.com:sous-chefs/rvm.git
git pull origin
commit 84ce9add8a421f278830ffb7192fc7d9b0e82438
Author: Lance Albertson <lance@osuosl.org>
Date:   Thu Oct 08 09:09:37 2020 -0700

    Merge pull request #406 from sous-chefs/automated/standardfiles

    Automated PR: Standardising Files
```
