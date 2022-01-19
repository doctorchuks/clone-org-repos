# Clone Organisation Repoistories

Clones or updates repositories from a GitHub organisation to local disk

## Prerequisites 

- Set a `GITHUB_TOKEN` environment variable to one which can pull from all the repositories you need

## Usage

clone-org-repos has 2 arguments:

- `org`, `-o`, Name of the org you wish to checkout, Required. 
- `path`, `-p`, Path to checkout repositories to, Optional, if not supplied it defaults to user's home directory


```bash
$ ./clone-org-repos -o sous-chefs

Cloning repository: rvm
SSH CLone URL: git@github.com:sous-chefs/rvm.git
git pull origin
commit 84ce9add8a421f278830ffb7192fc7d9b0e82438
Author: Lance Albertson <lance@osuosl.org>
Date:   Thu Oct 08 09:09:37 2020 -0700

    Merge pull request #406 from sous-chefs/automated/standardfiles

    Automated PR: Standardising Files
```

```bash
$ ./clone-org-repos -o sous-chefs -p ~/mydevfolder/sous-chefs

Cloning repository: rvm
SSH CLone URL: git@github.com:sous-chefs/rvm.git
git pull origin
commit 84ce9add8a421f278830ffb7192fc7d9b0e82438
Author: Lance Albertson <lance@osuosl.org>
Date:   Thu Oct 08 09:09:37 2020 -0700

    Merge pull request #406 from sous-chefs/automated/standardfiles

    Automated PR: Standardising Files
```
