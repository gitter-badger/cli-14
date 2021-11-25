# Installation
## Linux & Mac OS
### Global
<code>sh -c "$(curl -fsSL https://raw.githubusercontent.com/jarvis-technologies/cli/main/install.sh)"</code>
### Local
<code>
  curl -LO https://github.com/jarvis-technologies/cli/releases/download/v0.1.2/jt-os-here-arch-here
</code>

#### *Instructions for installing on Windows coming soon

# Usage
## Help
- <code>jt --help</code>
- <code>jt -h</code>

## Serve your application
- <code>jt ss</code>
- <code>jt run</code>
- <code>jt serve --port 8000 --folder ./public</code>
- <code>jt serve 8000 ./public</code>
### Options
- **--port** (_-p_) Port for PHP Dev Server (default is _8000_)
- **--folder** (_-f_) Public folder For PHP Dev Server (default is _./public_)

## Generate Http Controller
- <code>jt cc</code>
- <code>jt create:controller</code>
- <code>jt cc --name HomeController --path ./app/controllers/ --namespace JtF\\Application</code>
### Options
- **--name** (_-n_) Name of Controller (default is _HomeController_)
- **--namespace** Namespace of Controller (default is _JtF\\Application_)
- **--path** (_-p_) Path to generate `(!Must be directory)` (default is _./src/controllers/_) 

## Enable User Secrets for your application
- <code>jt secrets</code>

## Thanks for using JT Framework and other my products 😄
