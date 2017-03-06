func dumpJSON(v interface{}, kn string) {
  iterMap := func(x map[string]interface{}, root string) {
    var knf string
    if root == "root" {
      knf = "%q:%q"
    } else {
      knf = "%s:%q"
    }
    for k, v := range x {
      dumpJSON(v, fmt.Sprintf(knf, root, k))
    }
  }

  iterSlice := func(x []interface{}, root string) {
    var knf string
    if root == "root" {
      knf = "%q:[%d]"
    } else {
      knf = "%s:[%d]"
    }
    for k, v := range x {
      dumpJSON(v, fmt.Sprintf(knf, root, k))
    }
  }

  switch vv := v.(type) {
  case string:
    fmt.Printf("%s => (string) %q\n", kn, vv)
  case bool:
    fmt.Printf("%s => (bool) %v\n", kn, vv)
  case float64:
    fmt.Printf("%s => (float64) %f\n", kn, vv)
  case map[string]interface{}:
    fmt.Printf("%s => (map[string]interface{}) ...\n", kn)
    iterMap(vv, kn)
  case []interface{}:
    fmt.Printf("%s => ([]interface{}) ...\n", kn)
    iterSlice(vv, kn)
  default:
    fmt.Printf("%s => (unknown?) ...\n", kn)
  }
}