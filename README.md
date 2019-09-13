### Go client example for listing Pods

```
$ ./kube-cli podz
listing running Pods
**- Currently Running Pods -**
+-----------------------------------------------------------------+-------------+
|                            POD NAME                             |  NAMESPACE  |
+-----------------------------------------------------------------+-------------+
| ambassador-755657cbb-228q5                                      | default     |
| ambassador-consul-connect-integration-556bf7975b-6gd88          | default     |
| myconsul-consul-connect-injector-webhook-deployment-7579865zpw9 | default     |
| myconsul-consul-qtjbg                                           | default     |
| myconsul-consul-server-0                                        | default     |
| qotm-mtls-f769cb6ff-2m5t5                                       | default     |
| coredns-5c98db65d4-cz8kt                                        | kube-system |
| coredns-5c98db65d4-pqp4d                                        | kube-system |
| etcd-minikube                                                   | kube-system |
| heapster-gwhz6                                                  | kube-system |
| influxdb-grafana-7d9c2                                          | kube-system |
| kube-addon-manager-minikube                                     | kube-system |
| kube-apiserver-minikube                                         | kube-system |
| kube-controller-manager-minikube                                | kube-system |
| kube-proxy-qx2gk                                                | kube-system |
| kube-scheduler-minikube                                         | kube-system |
| kubernetes-dashboard-7b8ddcb5d6-s4hxh                           | kube-system |
| metrics-server-84bb785897-jcljm                                 | kube-system |
| nginx-ingress-controller-5d9cf9c69f-6z2lc                       | kube-system |
| storage-provisioner                                             | kube-system |
| tiller-deploy-84cccc96b5-qglm9                                  | kube-system |
+-----------------------------------------------------------------+-------------+

```
