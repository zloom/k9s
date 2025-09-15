<img src="https://raw.githubusercontent.com/derailed/k9s/master/assets/k9s_small.png" align="right" width="200" height="auto"/>

# Release v0.13.2

## Notes

Thank you to all that contributed with flushing out issues and enhancements for K9s! I'll try to mark some of these issues as fixed. But if you don't mind grab the latest rev and see if we're happier with some of the fixes! If you've filed an issue please help me verify and close. Your support, kindness and awesome suggestions to make K9s better is as ever very much noticed and appreciated!

Also if you dig this tool, please make some noise on social! [@kitesurfer](https://twitter.com/kitesurfer)

---

### XRay Reloaded. Part Duh!

<img src="https://raw.githubusercontent.com/derailed/k9s/master/assets/k9s_xray.png"/>

Found a waffle thin issue in the Beryllium(Be) core causing K9s xray vision to only operate on one eye ;)
Should be all betta' now...

The `xray` command now takes an **optional** third argument for the target namespace ie `:xray dp fred` will show the Xray view for deployments in the `fred` namespace.

Supported resources:

* Pods
* Deployments
* Services
* StatefulSets
* DaemonSets
* ReplicaSets

Still watch out for that overbite!! hence please proceed with caution...

---

## Resolved Bugs/Features

* [Issue #500](https://github.com/zloom/k9s/issues/500)

---

<img src="https://raw.githubusercontent.com/derailed/k9s/master/assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC. All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
