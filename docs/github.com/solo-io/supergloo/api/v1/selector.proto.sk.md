
---
title: "selector.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `supergloo.solo.io` 
##### Types:


- [PodSelector](#PodSelector)
- [LabelSelector](#LabelSelector)
- [UpstreamSelector](#UpstreamSelector)
- [NamespaceSelector](#NamespaceSelector)
  



##### Source File: [github.com/solo-io/supergloo/api/v1/selector.proto](https://github.com/solo-io/supergloo/blob/master/api/v1/selector.proto)





---
### <a name="PodSelector">PodSelector</a>

 
specifies a method by which to select pods
with in a mesh for the application of rules and policies

```yaml
"label_selector": .supergloo.solo.io.PodSelector.LabelSelector
"upstream_selector": .supergloo.solo.io.PodSelector.UpstreamSelector
"namespace_selector": .supergloo.solo.io.PodSelector.NamespaceSelector

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `label_selector` | [.supergloo.solo.io.PodSelector.LabelSelector](../selector.proto.sk#LabelSelector) | select pods by their labels |  |
| `upstream_selector` | [.supergloo.solo.io.PodSelector.UpstreamSelector](../selector.proto.sk#UpstreamSelector) | select pods by their corresponding upstreams |  |
| `namespace_selector` | [.supergloo.solo.io.PodSelector.NamespaceSelector](../selector.proto.sk#NamespaceSelector) | select all pods within one or more namespaces |  |




---
### <a name="LabelSelector">LabelSelector</a>

 
select pods by their labels

```yaml
"labels_to_match": map<string, string>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `labels_to_match` | `map<string, string>` |  |  |




---
### <a name="UpstreamSelector">UpstreamSelector</a>

 
select pods based on their services or subsets of services.
upstream CRDs will be created by discovery corresponding to
kubernetes services and the available subsets of those services

```yaml
"upstreams": []core.solo.io.ResourceRef

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `upstreams` | [[]core.solo.io.ResourceRef](../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) | apply the selector to one or more of their upstreams by adding their refs here |  |




---
### <a name="NamespaceSelector">NamespaceSelector</a>

 
select all pods in these namespaces

```yaml
"namespaces": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `namespaces` | `[]string` |  |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->