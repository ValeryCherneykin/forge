# Forge — Template Manager

Forge is a tool to quickly copy ready-made templates (like `Dockerfile`, `.gitignore`) from a template folder into your current project. No hassle—just fast and done.

## Why Use It?

- Tired of creating the same files for every project?  
- Sick of hunting down old templates and dragging them around?  
Forge fixes that: pick a template, and it’s in your project in seconds.

## Installation 🛠️

Want to give it a shot? Download and set it up in a couple of minutes!

### What You Need

- [Git](https://git-scm.com/downloads) — to grab the code.  
- [Go](https://golang.org/dl/) (1.24+) — to build it.  
- A terminal (Linux, macOS, Windows with Git Bash).

### How to Download and Install

#### Grab it from GitHub
Clone the repo:  
```bash
git clone https://github.com/ValeryCherneykin/forge.git
cd forge
```

#### Build and Install
Run:  
```bash
go install ./cmd/...
```
Forge will land in `~/go/bin`. Add it to your PATH if needed:  
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
Add that line to `~/.bashrc` or `~/.zshrc` to keep it permanent.  
Check it works:  
```bash
forge
```
If you see templates, you’re good to go!

#### Set Up Templates
Forge looks for templates in `~/.forge/templates`. Create it and drop your files there:  
```bash
mkdir -p ~/.forge/templates
cp /path/to/templates/* ~/.forge/templates/
```

## How to Use

### Run It
Just type:  
```bash
forge
```
Or point to your own template folder:  
```bash
forge -dir=/path/to/templates
```

### Controls

- `j` / `↓`: Move down the list.  
- `k` / `↑`: Move up the list.  
- `:`: Type to filter (e.g., `git` for `.gitignore`), `enter` to apply, `esc` to cancel.  
- `enter` / `p`: Copy the selected template to the current folder.  
- `q` / `ctrl+c`: Exit.

### Example

1. Run `forge`.  
2. Hit `:` and type `doc` to find `docker-compose.yml`.  
3. Press `enter` — the file’s copied.  
4. See: `Create file /path/to/docker-compose.yml` in green.
