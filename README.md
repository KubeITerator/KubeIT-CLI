# KubeIT - Command Line Interface (CLI)

Dedicated client for use with the [KubeIT backend](https://github.com/KubeITerator/KubeIT). 
This client allows users to preconfigure and use server sided schemes as templates for workflow creation.

## Contents
- [Description](#description)
- [Installation](#installation)
- [Usage](#features)
- [FAQ](#faq)
## Description

The KubeIT CLI is a dedicated interface method for the KubeIT Backend. 
It contains helper functions for communications with the backend
and allows automated access to in- and output data via S3.

## Installation

Download the lastest release for your desired operating system. All following examples reference the Linux binary kubeit, 
for other operating systems exchange `./kubeit` with the specific program (e.g. `./kubeit.exe` for windows)

#### Initial configuration:

Configure your backend connection with:

```
    ./kubeit configure
```

This will show a configuration dialogue that asks for your KubeIT URL and access token. 
The information is by default saved in `~/.kubeit/config.json`. Other paths can be specified via the `-c` flag.

## Usage

If more information is needed use the `-h` flag to access the help page with an overview for the chosen command.

#### Configure local scheme

To use KubeIT first a local scheme must be configured. 
Local schemes are variations of globally available backend schemes and can pre-configure backend-schemes with non-default parameters.
Local schemes can be configured, via a dialogue, with:

```
    ./kubeit configure -s
```

In normal operation only required arguments will be shown. If you want to change defaulted backend parameters locally, 
use the `-e` flag to enable expert configuration mode. (e.g. `./kubeit configure -s -e`)


#### Schedule local scheme to the cluster

The KubeIT CLI allows local schemes to be configured and executed in one step. A very basic example for this is:

```
    ./kubeit create workflow -s YOUR-LOCAL-SCHEME-NAME -i YOUR-INPUTFILE -o YOUR-OUTPUTFILE -w 
```

`-s`: specified your local scheme  
`-i`: Your input-file on the local file system (The default-template accepts protein FASTA files (.faa))  
`-o`: Automatically download results to the specified destination  
`-w`: Watch the workflow execution and wait for either workflow completion or failure  


### Input upload in advance

Input files can either be automatically upload via `-i` or uploaded beforehand. This can be done with:

```
    ./kubeit create S3 -f YOUR-FILE
```

If uploaded beforehand KubeIT will print an associated time-limited download URL automatically. This URL can be used for future workflows.

### Specify a parameter on scheduling

Sometime the needed parameters can differ. For this KubeIT can receive parameters in workflow creation. Parameters are specified with the `-p` flag followed by the parameter name and its desired value.
Example: `-p "input.inputdata=https://example.com/test"`

### Delete a workflow

Workflows can be deleted with:

```
    ./kubeit delete workflow -n WORKFLOW-NAME
```

### Delete a collection of workflows

Multiple workflows can be deleted by specifying a project group.

```
    ./kubeit delete workflow -g PROJECT-GROUP-NAME
```

### Get a workflow status

If a workflow is not scheduled with the `-w` flag, its status can be accessed once with:

```
    ./kubeit get workflow -n WORKFLOW-NAME
```

or for a complete project group:

```
    ./kubeit get workflow -g PROJECT-GROUP-NAME
```

### Get more information about (a) local or remote scheme(s)

Get all available schemes:

```
    ./kubeit get scheme
```

get more information about a specific scheme:

```
    ./kubeit get scheme -n SCHEME-NAME
```

both commands can either be used for remote schemes, or for local schemes (with the `-l` flag)

### Create new remote schemes

New remote schemes can be created with:

```
    ./kubeit create scheme -n SCHEME-NAME -f SCHEME-YAML-FILE
```

see [here](/docs/GUIDELINES.md) for a short guideline on how to create new schemes.


## FAQ

### How does KubeIT upload data to S3 ?

KubeIT uses pre-authenticated S3 links to upload data to S3. It automatically determines if a file have to be split up
or can be uploaded as a whole.

### How does the splitting work ?

Splitting is done using the Splitter WorkflowTemplate and the splitter interface. For mor information see [here]()
