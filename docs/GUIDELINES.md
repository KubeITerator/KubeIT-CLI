# Guidelines for the creation of new Schemes

This is a small collection of guidelines for creating new Schemes in KubeIT.
KubeIT schemes are modified Argo workflows. For a good start get familar with Argo workflow syntax [here](https://argoproj.github.io/argo-workflows/examples/)
You can use any of Argos features in your scheme. A basic example for a scheme as reference is found [here](https://github.com/KubeITerator/KubeIT/blob/master/default-settings/default-template.yaml)

## How to create a scheme

1. Start by creating a regular Argo workflow. 

2. Determine arguments that need to be dynamic and should be interchangeable via parameters.

3. Use a similar metadata structure as this:

```
  generateName: {{kubeit.workflow.nameprefix}} # Default: "kubeit-test"
  namespace: {{kubeit.workflow.namespace}}
  labels:
    project: {{kubeit.project.name}}
```
This metadata section enables you to use the full potential of KubeIT. generateName is optional and can have a fixed value.
The label "project" is used by kubeIT to group multiple workflows together, if this label is omitted the project or group feature will not work.

4. Create your KubeIT parameters

KubeIT parameters look similar to Argo parameters, like this `{{kubeit.category.name}}`. Category and name can have any value.
Specified KubeIT parameters are (by default) required. They need to be specified to the KubeIT API for workflow scheduling to occur.
If you want to se a default value for a kubeit parameter: add a comment with `# Default: "DEFAULT-VALUE""`. The keyword Default is recognized by KubeIT
and everything in the following quotes is used as default value for the previous parameter.

5. Use WorkflowTemplates for your actual workload
   An overview of available default WorkflowTemplates in KubeIT can be found [here](). WorkflowTemplates allow your schemes to be vastly more modular.

5. Splitting

The overall splitting design is described [here](). Splitting uses the built-in argo feature `withItems` that multiplexes a single JSON list to multiple executed Tasks.