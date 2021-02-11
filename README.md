# KubeIT - Command Line Interface (CLI)

Dedicated client for use with the [KubeIT backend](https://github.com/KubeITerator/KubeIT). 
This client allows users to preconfigure and use server sided schemes as templates for workflow creation.

## Contents
- [Installation](#installation)
- [Usage](#features)
- [Availability](#availability)
- [Input & Output](#input-output)
- [FAQ](#faq)
- [License](#license)
- [Bugs](#bugs)


## Installation

Download client-release artifact and unzip it.
Set the needed environment variables with:
```
    export K8SEDGAR_URL="YOUR-BIOKUBE-URL"
```
#### Initial configuration:

- K8SEDGAR_URL: Biokube-URL (with https prefix)
- K8SEDGAR_TOKEN: X-Auth-Token for k8s server.
- optional: K8SEDGAR_BBUCKET: S3 Basepath for File-Upload (default: edgartest)


## Usage

Scheduling a single job to kubernetes:

```
    ./k8sedgar create -i {INPUTFILE} -o {OUTPUTFILE}
```

INPUTFILE expected format: fasta (protein)
OUTPUTFILE format: Plain text

optional parameters are:

- -s: Chunksize in Bytes for distributing Blast-Jobs (default: 1000000)
- -d: Daemonize the job and wait for results in an independent process.

If using the daemon process it can be terminated with:

```
    ./k8sedgar daemon stop
```

And the status of not finished jobs can be determined with:

```
    ./k8sedgar status
```





## Usage for multi-jobs with config-file

### Scheduling
Scheduling a job requires a json config file. To configure your job just create a json-config
file with the following scheme. The "JobParameter" and "SplitSize" (default: 10 MB) parameter are optional and can be
omitted.

```
{

  "JobParameter": [ PARAMETERNAME=PARAMETERVALUE ],
  "OutputFolder": "FOLDER FOR RESULTS",
  "SplitSize":     INTEGER SIZE (MB) FOR SPLITTING,
  "InputFiles": [
                  "PATH TO INPUTFILE",
                  "PATH TO ANOTHER INPUTFILE"
                ]
}
```

### Run the Job

To run the job just use:
```
    ./k8sedgar create -f job.json
```

to run the job and detach a daemon process for monitoring run:

```
    ./k8sedgar create -f job.json -d
```

### Get status

To get the status of currently monitored jobs just run:
```
    ./k8sedgar status
```

### Terminate background daemon

If something went wrong, and the background process behaves not properly.
The background daemon process can be terminated with:

```
    ./k8sedgar daemon stop
```


    
