
# 一些基础命令

查看 Kubernetes 支持的资源类型：
```shell
kubectl api-resources
```

运行结果如下
```shell
kubectl api-resources
NAME                                SHORTNAMES   APIVERSION                        NAMESPACED   KIND
bindings                                         v1                                true         Binding
componentstatuses                   cs           v1                                false        ComponentStatus
configmaps                          cm           v1                                true         ConfigMap
endpoints                           ep           v1                                true         Endpoints
events                              ev           v1                                true         Event
limitranges                         limits       v1                                true         LimitRange
namespaces                          ns           v1                                false        Namespace
nodes                               no           v1                                false        Node
persistentvolumeclaims              pvc          v1                                true         PersistentVolumeClaim
persistentvolumes                   pv           v1                                false        PersistentVolume
pods                                po           v1                                true         Pod
podtemplates                                     v1                                true         PodTemplate
replicationcontrollers              rc           v1                                true         ReplicationController
resourcequotas                      quota        v1                                true         ResourceQuota
secrets                                          v1                                true         Secret
serviceaccounts                     sa           v1                                true         ServiceAccount
services                            svc          v1                                true         Service
mutatingwebhookconfigurations                    admissionregistration.k8s.io/v1   false        MutatingWebhookConfiguration
validatingadmissionpolicies                      admissionregistration.k8s.io/v1   false        ValidatingAdmissionPolicy
validatingadmissionpolicybindings                admissionregistration.k8s.io/v1   false        ValidatingAdmissionPolicyBinding
validatingwebhookconfigurations                  admissionregistration.k8s.io/v1   false        ValidatingWebhookConfiguration
customresourcedefinitions           crd,crds     apiextensions.k8s.io/v1           false        CustomResourceDefinition
apiservices                                      apiregistration.k8s.io/v1         false        APIService
controllerrevisions                              apps/v1                           true         ControllerRevision
daemonsets                          ds           apps/v1                           true         DaemonSet
deployments                         deploy       apps/v1                           true         Deployment
replicasets                         rs           apps/v1                           true         ReplicaSet
statefulsets                        sts          apps/v1                           true         StatefulSet
selfsubjectreviews                               authentication.k8s.io/v1          false        SelfSubjectReview
tokenreviews                                     authentication.k8s.io/v1          false        TokenReview
localsubjectaccessreviews                        authorization.k8s.io/v1           true         LocalSubjectAccessReview
selfsubjectaccessreviews                         authorization.k8s.io/v1           false        SelfSubjectAccessReview
selfsubjectrulesreviews                          authorization.k8s.io/v1           false        SelfSubjectRulesReview
subjectaccessreviews                             authorization.k8s.io/v1           false        SubjectAccessReview
horizontalpodautoscalers            hpa          autoscaling/v2                    true         HorizontalPodAutoscaler
cronjobs                            cj           batch/v1                          true         CronJob
jobs                                             batch/v1                          true         Job
certificatesigningrequests          csr          certificates.k8s.io/v1            false        CertificateSigningRequest
leases                                           coordination.k8s.io/v1            true         Lease
bgpconfigurations                                crd.projectcalico.org/v1          false        BGPConfiguration
bgpfilters                                       crd.projectcalico.org/v1          false        BGPFilter
bgppeers                                         crd.projectcalico.org/v1          false        BGPPeer
blockaffinities                                  crd.projectcalico.org/v1          false        BlockAffinity
caliconodestatuses                               crd.projectcalico.org/v1          false        CalicoNodeStatus
clusterinformations                              crd.projectcalico.org/v1          false        ClusterInformation
felixconfigurations                              crd.projectcalico.org/v1          false        FelixConfiguration
globalnetworkpolicies                            crd.projectcalico.org/v1          false        GlobalNetworkPolicy
globalnetworksets                                crd.projectcalico.org/v1          false        GlobalNetworkSet
hostendpoints                                    crd.projectcalico.org/v1          false        HostEndpoint
ipamblocks                                       crd.projectcalico.org/v1          false        IPAMBlock
ipamconfigs                                      crd.projectcalico.org/v1          false        IPAMConfig
ipamhandles                                      crd.projectcalico.org/v1          false        IPAMHandle
ippools                                          crd.projectcalico.org/v1          false        IPPool
ipreservations                                   crd.projectcalico.org/v1          false        IPReservation
kubecontrollersconfigurations                    crd.projectcalico.org/v1          false        KubeControllersConfiguration
networkpolicies                                  crd.projectcalico.org/v1          true         NetworkPolicy
networksets                                      crd.projectcalico.org/v1          true         NetworkSet
endpointslices                                   discovery.k8s.io/v1               true         EndpointSlice
events                              ev           events.k8s.io/v1                  true         Event
flowschemas                                      flowcontrol.apiserver.k8s.io/v1   false        FlowSchema
prioritylevelconfigurations                      flowcontrol.apiserver.k8s.io/v1   false        PriorityLevelConfiguration
ingressclasses                                   networking.k8s.io/v1              false        IngressClass
ingresses                           ing          networking.k8s.io/v1              true         Ingress
networkpolicies                     netpol       networking.k8s.io/v1              true         NetworkPolicy
runtimeclasses                                   node.k8s.io/v1                    false        RuntimeClass
poddisruptionbudgets                pdb          policy/v1                         true         PodDisruptionBudget
clusterrolebindings                              rbac.authorization.k8s.io/v1      false        ClusterRoleBinding
clusterroles                                     rbac.authorization.k8s.io/v1      false        ClusterRole
rolebindings                                     rbac.authorization.k8s.io/v1      true         RoleBinding
roles                                            rbac.authorization.k8s.io/v1      true         Role
priorityclasses                     pc           scheduling.k8s.io/v1              false        PriorityClass
csidrivers                                       storage.k8s.io/v1                 false        CSIDriver
csinodes                                         storage.k8s.io/v1                 false        CSINode
csistoragecapacities                             storage.k8s.io/v1                 true         CSIStorageCapacity
storageclasses                      sc           storage.k8s.io/v1                 false        StorageClass
volumeattachments                                storage.k8s.io/v1                 false        VolumeAttachment
```

在输出的`“NAME”`一栏，就是对象的名字，比如 `ConfigMap`、`Pod`、`Service` 等等，第二栏`“SHORTNAMES”`则是这种资源的简写，在我们使用 `kubectl` 命令的时候很有用，可以少敲几次键盘，比如 `Pod` 可以简写成 `po`，`Service` 可以简写成 `svc`


在使用 `kubectl` 命令的时候，你还可以加上一个参数 `--v=9`，它会显示出详细的命令执行过程，清楚地看到发出的 `HTTP` 请求，比如：
```shell
$ kubectl get pod --v=9
I0330 08:34:25.072478  337675 loader.go:395] Config loaded from file:  /root/.kube/config
I0330 08:34:25.076051  337675 round_trippers.go:466] curl -v -XGET  -H "Accept: application/json;as=Table;v=v1;g=meta.k8s.io,application/json;as=Table;v=v1beta1;g=meta.k8s.io,application/json" -H "User-Agent: kubectl/v1.30.14 (linux/amd64) kubernetes/9e18483" 'https://192.168.44.130:6443/api/v1/namespaces/default/pods?limit=500'
I0330 08:34:25.076462  337675 round_trippers.go:510] HTTP Trace: Dial to tcp:192.168.44.130:6443 succeed
I0330 08:34:25.080362  337675 round_trippers.go:553] GET https://192.168.44.130:6443/api/v1/namespaces/default/pods?limit=500 200 OK in 4 milliseconds
I0330 08:34:25.080376  337675 round_trippers.go:570] HTTP Statistics: DNSLookup 0 ms Dial 0 ms TLSHandshake 2 ms ServerProcessing 1 ms Duration 4 ms
I0330 08:34:25.080380  337675 round_trippers.go:577] Response Headers:
I0330 08:34:25.080384  337675 round_trippers.go:580]     Content-Length: 2961
I0330 08:34:25.080387  337675 round_trippers.go:580]     Date: Mon, 30 Mar 2026 00:34:25 GMT
I0330 08:34:25.080389  337675 round_trippers.go:580]     Audit-Id: 313796da-23a2-426b-97c9-54737ccd4bfd
I0330 08:34:25.080392  337675 round_trippers.go:580]     Cache-Control: no-cache, private
I0330 08:34:25.080397  337675 round_trippers.go:580]     Content-Type: application/json
I0330 08:34:25.080399  337675 round_trippers.go:580]     X-Kubernetes-Pf-Flowschema-Uid: 9fe55451-5a66-4b86-b498-f85ec51121ea
I0330 08:34:25.080401  337675 round_trippers.go:580]     X-Kubernetes-Pf-Prioritylevel-Uid: 24a610da-be1f-49b0-81d9-7b799dfc77d4
I0330 08:34:25.080482  337675 request.go:1212] Response Body: {"kind":"Table","apiVersion":"meta.k8s.io/v1","metadata":{"resourceVersion":"446677"},"columnDefinitions":[{"name":"Name","type":"string","format":"name","description":"Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names","priority":0},{"name":"Ready","type":"string","format":"","description":"The aggregate readiness state of this pod for accepting traffic.","priority":0},{"name":"Status","type":"string","format":"","description":"The aggregate status of the containers in this pod.","priority":0},{"name":"Restarts","type":"string","format":"","description":"The number of times the containers in this pod have been restarted and when the last container in this pod has restarted.","priority":0},{"name":"Age","type":"string","format":"","description":"CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata","priority":0},{"name":"IP","type":"string","format":"","description":"podIP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.","priority":1},{"name":"Node","type":"string","format":"","description":"NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.","priority":1},{"name":"Nominated Node","type":"string","format":"","description":"nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.","priority":1},{"name":"Readiness Gates","type":"string","format":"","description":"If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to \"True\" More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates","priority":1}],"rows":[]}
```