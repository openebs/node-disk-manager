# Git Development Workflow

## Initial Setup

### Fork in the cloud

1. Visit https://github.com/openebs/node-disk-manager
2. Click `Fork` button (top right) to establish a cloud-based fork.

### Clone fork to local host

Since Go modules are being used, you can clone the repo in any directory on your local system. 
Create your clone:

```sh

# Note: Here user= your github profile name
git clone https://github.com/$user/node-disk-manager.git

# Configure remote upstream
cd node-disk-manager
git remote add upstream https://github.com/openebs/node-disk-manager.git

# Never push to upstream develop
git remote set-url --push upstream no_push

# Confirm that your remotes make sense:
git remote -v
```


### Always sync your local repository:
Open a terminal on your local host. Change directory to the node-disk-manager fork root.

```sh
$ cd node-disk-manager
```

 Checkout the develop branch.

 ```sh
 $ git checkout develop
 Switched to branch 'develop'
 Your branch is up-to-date with 'origin/develop'.
 ```

 Recall that origin/develop is a branch on your remote GitHub repository.
 Make sure you have the upstream remote openebs/node-disk-manager by listing them.

 ```sh
 $ git remote -v
 origin	https://github.com/$user/node-disk-manager.git (fetch)
 origin	https://github.com/$user/node-disk-manager.git (push)
 upstream	https://github.com/openebs/node-disk-manager.git (fetch)
 upstream	https://github.com/openebs/node-disk-manager.git (no_push)
 ```

 If the upstream is missing, add it by using below command.

 ```sh
 $ git remote add upstream https://github.com/openebs/node-disk-manager.git
 ```
 Fetch all the changes from the upstream develop branch.

 ```sh
 $ git fetch upstream develop
 remote: Counting objects: 141, done.
 remote: Compressing objects: 100% (29/29), done.
 remote: Total 141 (delta 52), reused 46 (delta 46), pack-reused 66
 Receiving objects: 100% (141/141), 112.43 KiB | 0 bytes/s, done.
 Resolving deltas: 100% (79/79), done.
 From github.com:openebs/node-disk-manager
   * branch            develop     -> FETCH_HEAD
 ```

 Rebase your local develop with the upstream/develop.

 ```sh
 $ git rebase upstream/develop
 First, rewinding head to replay your work on top of it...
 Fast-forwarded develop to upstream/develop.
 ```
 This command applies all the commits from the upstream develop to your local develop.

 Check the status of your local branch.

 ```sh
 $ git status
 On branch develop
 Your branch is ahead of 'origin/develop' by 38 commits.
 (use "git push" to publish your local commits)
 nothing to commit, working directory clean
 ```
 Your local repository now has all the changes from the upstream remote. You need to push the changes to your own remote fork which is origin develop.

 Push the rebased develop to origin develop.

 ```sh
 $ git push origin develop
 Username for 'https://github.com': $user
 Password for 'https://$user@github.com':
 Counting objects: 223, done.
 Compressing objects: 100% (38/38), done.
 Writing objects: 100% (69/69), 8.76 KiB | 0 bytes/s, done.
 Total 69 (delta 53), reused 47 (delta 31)
 To https://github.com/$user/node-disk-manager.git
 8e107a9..5035fa1  develop -> develop
 ```

### Contributing to a feature or bugfix. 

Always start with creating a new branch from develop to work on a new feature or bugfix. Your branch name should have the format XX-descriptive where XX is the issue number you are working on followed by some descriptive text. For example:

 ```sh
 $ git checkout develop
 # Make sure the develop is rebased with the latest changes as described in previous step.
 $ git checkout -b 1234-fix-developer-docs
 Switched to a new branch '1234-fix-developer-docs'
 ```
Happy Hacking!

### Keep your branch in sync

[Rebasing](https://git-scm.com/docs/git-rebase) is very import to keep your branch in sync with the changes being made by others and to avoid huge merge conflicts while raising your Pull Requests. You will always have to rebase before raising the PR. 

```sh
# While on your myfeature branch (see above)
git fetch upstream
git rebase upstream/develop
```

While you rebase your changes, you must resolve any conflicts that might arise and build and test your changes using the above steps. 

## Submission

### Create a pull request

Before you raise the Pull Requests, ensure you have reviewed the checklist in the [CONTRIBUTING GUIDE](https://github.com/openebs/openebs/blob/main/CONTRIBUTING.md):
- Ensure that you have re-based your changes with the upstream using the steps above.
- Ensure that you have added the required unit tests for the bug fixes or new feature that you have introduced.
- Ensure that commits are signed (DCO) .  
- Ensure your commits history is clean with proper header and descriptions.

Go to the [openebs/node-disk-manager github](https://github.com/openebs/node-disk-manager) and follow the Open Pull Request link to raise your PR from your development branch.

