/*
Copyright 2023 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes

const (
	apiDiscovery = `{"kind":"APIGroupDiscoveryList","apiVersion":"apidiscovery.k8s.io/v2beta1","metadata":{},"items":[{"metadata":{"creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"bindings","responseKind":{"group":"","version":"","kind":"Binding"},"scope":"Namespaced","singularResource":"binding","verbs":["create"]},{"resource":"componentstatuses","responseKind":{"group":"","version":"","kind":"ComponentStatus"},"scope":"Cluster","singularResource":"componentstatus","verbs":["get","list"],"shortNames":["cs"]},{"resource":"configmaps","responseKind":{"group":"","version":"","kind":"ConfigMap"},"scope":"Namespaced","singularResource":"configmap","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["cm"]},{"resource":"endpoints","responseKind":{"group":"","version":"","kind":"Endpoints"},"scope":"Namespaced","singularResource":"endpoints","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["ep"]},{"resource":"events","responseKind":{"group":"","version":"","kind":"Event"},"scope":"Namespaced","singularResource":"event","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["ev"]},{"resource":"limitranges","responseKind":{"group":"","version":"","kind":"LimitRange"},"scope":"Namespaced","singularResource":"limitrange","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["limits"]},{"resource":"namespaces","responseKind":{"group":"","version":"","kind":"Namespace"},"scope":"Cluster","singularResource":"namespace","verbs":["create","delete","get","list","patch","update","watch"],"shortNames":["ns"],"subresources":[{"subresource":"finalize","responseKind":{"group":"","version":"","kind":"Namespace"},"verbs":["update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"Namespace"},"verbs":["get","patch","update"]}]},{"resource":"nodes","responseKind":{"group":"","version":"","kind":"Node"},"scope":"Cluster","singularResource":"node","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["no"],"subresources":[{"subresource":"proxy","responseKind":{"group":"","version":"","kind":"NodeProxyOptions"},"verbs":["create","delete","get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"Node"},"verbs":["get","patch","update"]}]},{"resource":"persistentvolumeclaims","responseKind":{"group":"","version":"","kind":"PersistentVolumeClaim"},"scope":"Namespaced","singularResource":"persistentvolumeclaim","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["pvc"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"PersistentVolumeClaim"},"verbs":["get","patch","update"]}]},{"resource":"persistentvolumes","responseKind":{"group":"","version":"","kind":"PersistentVolume"},"scope":"Cluster","singularResource":"persistentvolume","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["pv"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"PersistentVolume"},"verbs":["get","patch","update"]}]},{"resource":"pods","responseKind":{"group":"","version":"","kind":"Pod"},"scope":"Namespaced","singularResource":"pod","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["po"],"categories":["all"],"subresources":[{"subresource":"attach","responseKind":{"group":"","version":"","kind":"PodAttachOptions"},"verbs":["create","get"]},{"subresource":"binding","responseKind":{"group":"","version":"","kind":"Binding"},"verbs":["create"]},{"subresource":"ephemeralcontainers","responseKind":{"group":"","version":"","kind":"Pod"},"verbs":["get","patch","update"]},{"subresource":"eviction","responseKind":{"group":"policy","version":"v1","kind":"Eviction"},"verbs":["create"]},{"subresource":"exec","responseKind":{"group":"","version":"","kind":"PodExecOptions"},"verbs":["create","get"]},{"subresource":"log","responseKind":{"group":"","version":"","kind":"Pod"},"verbs":["get"]},{"subresource":"portforward","responseKind":{"group":"","version":"","kind":"PodPortForwardOptions"},"verbs":["create","get"]},{"subresource":"proxy","responseKind":{"group":"","version":"","kind":"PodProxyOptions"},"verbs":["create","delete","get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"Pod"},"verbs":["get","patch","update"]}]},{"resource":"podtemplates","responseKind":{"group":"","version":"","kind":"PodTemplate"},"scope":"Namespaced","singularResource":"podtemplate","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"replicationcontrollers","responseKind":{"group":"","version":"","kind":"ReplicationController"},"scope":"Namespaced","singularResource":"replicationcontroller","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["rc"],"categories":["all"],"subresources":[{"subresource":"scale","responseKind":{"group":"autoscaling","version":"v1","kind":"Scale"},"verbs":["get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"ReplicationController"},"verbs":["get","patch","update"]}]},{"resource":"resourcequotas","responseKind":{"group":"","version":"","kind":"ResourceQuota"},"scope":"Namespaced","singularResource":"resourcequota","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["quota"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"ResourceQuota"},"verbs":["get","patch","update"]}]},{"resource":"secrets","responseKind":{"group":"","version":"","kind":"Secret"},"scope":"Namespaced","singularResource":"secret","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"serviceaccounts","responseKind":{"group":"","version":"","kind":"ServiceAccount"},"scope":"Namespaced","singularResource":"serviceaccount","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["sa"],"subresources":[{"subresource":"token","responseKind":{"group":"authentication.k8s.io","version":"v1","kind":"TokenRequest"},"verbs":["create"]}]},{"resource":"services","responseKind":{"group":"","version":"","kind":"Service"},"scope":"Namespaced","singularResource":"service","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["svc"],"categories":["all"],"subresources":[{"subresource":"proxy","responseKind":{"group":"","version":"","kind":"ServiceProxyOptions"},"verbs":["create","delete","get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"Service"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]}]}`

	apisDiscovery = `{"kind":"APIGroupDiscoveryList","apiVersion":"apidiscovery.k8s.io/v2beta1","metadata":{},"items":[{"metadata":{"name":"apiregistration.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"apiservices","responseKind":{"group":"","version":"","kind":"APIService"},"scope":"Cluster","singularResource":"apiservice","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"categories":["api-extensions"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"APIService"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"apps","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"controllerrevisions","responseKind":{"group":"","version":"","kind":"ControllerRevision"},"scope":"Namespaced","singularResource":"controllerrevision","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"daemonsets","responseKind":{"group":"","version":"","kind":"DaemonSet"},"scope":"Namespaced","singularResource":"daemonset","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["ds"],"categories":["all"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"DaemonSet"},"verbs":["get","patch","update"]}]},{"resource":"deployments","responseKind":{"group":"","version":"","kind":"Deployment"},"scope":"Namespaced","singularResource":"deployment","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["deploy"],"categories":["all"],"subresources":[{"subresource":"scale","responseKind":{"group":"autoscaling","version":"v1","kind":"Scale"},"verbs":["get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"Deployment"},"verbs":["get","patch","update"]}]},{"resource":"replicasets","responseKind":{"group":"","version":"","kind":"ReplicaSet"},"scope":"Namespaced","singularResource":"replicaset","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["rs"],"categories":["all"],"subresources":[{"subresource":"scale","responseKind":{"group":"autoscaling","version":"v1","kind":"Scale"},"verbs":["get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"ReplicaSet"},"verbs":["get","patch","update"]}]},{"resource":"statefulsets","responseKind":{"group":"","version":"","kind":"StatefulSet"},"scope":"Namespaced","singularResource":"statefulset","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["sts"],"categories":["all"],"subresources":[{"subresource":"scale","responseKind":{"group":"autoscaling","version":"v1","kind":"Scale"},"verbs":["get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"StatefulSet"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"events.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"events","responseKind":{"group":"","version":"","kind":"Event"},"scope":"Namespaced","singularResource":"event","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["ev"]}],"freshness":"Current"}]},{"metadata":{"name":"authentication.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"tokenreviews","responseKind":{"group":"","version":"","kind":"TokenReview"},"scope":"Cluster","singularResource":"tokenreview","verbs":["create"]}],"freshness":"Current"}]},{"metadata":{"name":"authorization.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"localsubjectaccessreviews","responseKind":{"group":"","version":"","kind":"LocalSubjectAccessReview"},"scope":"Namespaced","singularResource":"localsubjectaccessreview","verbs":["create"]},{"resource":"selfsubjectaccessreviews","responseKind":{"group":"","version":"","kind":"SelfSubjectAccessReview"},"scope":"Cluster","singularResource":"selfsubjectaccessreview","verbs":["create"]},{"resource":"selfsubjectrulesreviews","responseKind":{"group":"","version":"","kind":"SelfSubjectRulesReview"},"scope":"Cluster","singularResource":"selfsubjectrulesreview","verbs":["create"]},{"resource":"subjectaccessreviews","responseKind":{"group":"","version":"","kind":"SubjectAccessReview"},"scope":"Cluster","singularResource":"subjectaccessreview","verbs":["create"]}],"freshness":"Current"}]},{"metadata":{"name":"autoscaling","creationTimestamp":null},"versions":[{"version":"v2","resources":[{"resource":"horizontalpodautoscalers","responseKind":{"group":"","version":"","kind":"HorizontalPodAutoscaler"},"scope":"Namespaced","singularResource":"horizontalpodautoscaler","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["hpa"],"categories":["all"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"HorizontalPodAutoscaler"},"verbs":["get","patch","update"]}]}],"freshness":"Current"},{"version":"v1","resources":[{"resource":"horizontalpodautoscalers","responseKind":{"group":"","version":"","kind":"HorizontalPodAutoscaler"},"scope":"Namespaced","singularResource":"horizontalpodautoscaler","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["hpa"],"categories":["all"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"HorizontalPodAutoscaler"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"batch","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"cronjobs","responseKind":{"group":"","version":"","kind":"CronJob"},"scope":"Namespaced","singularResource":"cronjob","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["cj"],"categories":["all"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"CronJob"},"verbs":["get","patch","update"]}]},{"resource":"jobs","responseKind":{"group":"","version":"","kind":"Job"},"scope":"Namespaced","singularResource":"job","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"categories":["all"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"Job"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"certificates.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"certificatesigningrequests","responseKind":{"group":"","version":"","kind":"CertificateSigningRequest"},"scope":"Cluster","singularResource":"certificatesigningrequest","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["csr"],"subresources":[{"subresource":"approval","responseKind":{"group":"","version":"","kind":"CertificateSigningRequest"},"verbs":["get","patch","update"]},{"subresource":"status","responseKind":{"group":"","version":"","kind":"CertificateSigningRequest"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"networking.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"ingressclasses","responseKind":{"group":"","version":"","kind":"IngressClass"},"scope":"Cluster","singularResource":"ingressclass","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"ingresses","responseKind":{"group":"","version":"","kind":"Ingress"},"scope":"Namespaced","singularResource":"ingress","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["ing"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"Ingress"},"verbs":["get","patch","update"]}]},{"resource":"networkpolicies","responseKind":{"group":"","version":"","kind":"NetworkPolicy"},"scope":"Namespaced","singularResource":"networkpolicy","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["netpol"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"NetworkPolicy"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"policy","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"poddisruptionbudgets","responseKind":{"group":"","version":"","kind":"PodDisruptionBudget"},"scope":"Namespaced","singularResource":"poddisruptionbudget","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["pdb"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"PodDisruptionBudget"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"rbac.authorization.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"clusterrolebindings","responseKind":{"group":"","version":"","kind":"ClusterRoleBinding"},"scope":"Cluster","singularResource":"clusterrolebinding","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"clusterroles","responseKind":{"group":"","version":"","kind":"ClusterRole"},"scope":"Cluster","singularResource":"clusterrole","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"rolebindings","responseKind":{"group":"","version":"","kind":"RoleBinding"},"scope":"Namespaced","singularResource":"rolebinding","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"roles","responseKind":{"group":"","version":"","kind":"Role"},"scope":"Namespaced","singularResource":"role","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]}],"freshness":"Current"}]},{"metadata":{"name":"storage.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"csidrivers","responseKind":{"group":"","version":"","kind":"CSIDriver"},"scope":"Cluster","singularResource":"csidriver","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"csinodes","responseKind":{"group":"","version":"","kind":"CSINode"},"scope":"Cluster","singularResource":"csinode","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"csistoragecapacities","responseKind":{"group":"","version":"","kind":"CSIStorageCapacity"},"scope":"Namespaced","singularResource":"csistoragecapacity","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"resource":"storageclasses","responseKind":{"group":"","version":"","kind":"StorageClass"},"scope":"Cluster","singularResource":"storageclass","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["sc"]},{"resource":"volumeattachments","responseKind":{"group":"","version":"","kind":"VolumeAttachment"},"scope":"Cluster","singularResource":"volumeattachment","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"VolumeAttachment"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"admissionregistration.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"mutatingwebhookconfigurations","responseKind":{"group":"","version":"","kind":"MutatingWebhookConfiguration"},"scope":"Cluster","singularResource":"mutatingwebhookconfiguration","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"categories":["api-extensions"]},{"resource":"validatingwebhookconfigurations","responseKind":{"group":"","version":"","kind":"ValidatingWebhookConfiguration"},"scope":"Cluster","singularResource":"validatingwebhookconfiguration","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"categories":["api-extensions"]}],"freshness":"Current"}]},{"metadata":{"name":"apiextensions.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"customresourcedefinitions","responseKind":{"group":"","version":"","kind":"CustomResourceDefinition"},"scope":"Cluster","singularResource":"customresourcedefinition","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["crd","crds"],"categories":["api-extensions"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"CustomResourceDefinition"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"scheduling.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"priorityclasses","responseKind":{"group":"","version":"","kind":"PriorityClass"},"scope":"Cluster","singularResource":"priorityclass","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["pc"]}],"freshness":"Current"}]},{"metadata":{"name":"coordination.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"leases","responseKind":{"group":"","version":"","kind":"Lease"},"scope":"Namespaced","singularResource":"lease","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]}],"freshness":"Current"}]},{"metadata":{"name":"node.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"runtimeclasses","responseKind":{"group":"","version":"","kind":"RuntimeClass"},"scope":"Cluster","singularResource":"runtimeclass","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]}],"freshness":"Current"}]},{"metadata":{"name":"discovery.k8s.io","creationTimestamp":null},"versions":[{"version":"v1","resources":[{"resource":"endpointslices","responseKind":{"group":"","version":"","kind":"EndpointSlice"},"scope":"Namespaced","singularResource":"endpointslice","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]}],"freshness":"Current"}]},{"metadata":{"name":"flowcontrol.apiserver.k8s.io","creationTimestamp":null},"versions":[{"version":"v1beta3","resources":[{"resource":"flowschemas","responseKind":{"group":"","version":"","kind":"FlowSchema"},"scope":"Cluster","singularResource":"flowschema","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"FlowSchema"},"verbs":["get","patch","update"]}]},{"resource":"prioritylevelconfigurations","responseKind":{"group":"","version":"","kind":"PriorityLevelConfiguration"},"scope":"Cluster","singularResource":"prioritylevelconfiguration","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"PriorityLevelConfiguration"},"verbs":["get","patch","update"]}]}],"freshness":"Current"},{"version":"v1beta2","resources":[{"resource":"flowschemas","responseKind":{"group":"","version":"","kind":"FlowSchema"},"scope":"Cluster","singularResource":"flowschema","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"FlowSchema"},"verbs":["get","patch","update"]}]},{"resource":"prioritylevelconfigurations","responseKind":{"group":"","version":"","kind":"PriorityLevelConfiguration"},"scope":"Cluster","singularResource":"prioritylevelconfiguration","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"subresources":[{"subresource":"status","responseKind":{"group":"","version":"","kind":"PriorityLevelConfiguration"},"verbs":["get","patch","update"]}]}],"freshness":"Current"}]},{"metadata":{"name":"dapr.io","creationTimestamp":null},"versions":[{"version":"v2alpha1","resources":[{"resource":"subscriptions","responseKind":{"group":"dapr.io","version":"v2alpha1","kind":"Subscription"},"scope":"Namespaced","singularResource":"subscription","verbs":["delete","deletecollection","get","list","patch","create","update","watch"],"categories":["all","dapr"]}],"freshness":"Current"},{"version":"v1alpha1","resources":[{"resource":"components","responseKind":{"group":"dapr.io","version":"v1alpha1","kind":"Component"},"scope":"Namespaced","singularResource":"component","verbs":["delete","deletecollection","get","list","patch","create","update","watch"],"categories":["all","dapr"]},{"resource":"configurations","responseKind":{"group":"dapr.io","version":"v1alpha1","kind":"Configuration"},"scope":"Namespaced","singularResource":"configuration","verbs":["delete","deletecollection","get","list","patch","create","update","watch"]},{"resource":"httpendpoints","responseKind":{"group":"dapr.io","version":"v1alpha1","kind":"HTTPEndpoint"},"scope":"Namespaced","singularResource":"httpendpoint","verbs":["delete","deletecollection","get","list","patch","create","update","watch"]},{"resource":"resiliencies","responseKind":{"group":"dapr.io","version":"v1alpha1","kind":"Resiliency"},"scope":"Namespaced","singularResource":"resiliency","verbs":["delete","deletecollection","get","list","patch","create","update","watch"],"categories":["dapr"]},{"resource":"subscriptions","responseKind":{"group":"dapr.io","version":"v1alpha1","kind":"Subscription"},"scope":"Namespaced","singularResource":"subscription","verbs":["delete","deletecollection","get","list","patch","create","update","watch"],"categories":["all","dapr"]}],"freshness":"Current"}]}]}`
)