package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"ldif", }, `{"contains":[{"className":"attribute","begin":"^dn","end":": ","excludeEnd":true,"starts":{"end":"$","relevance":0},"relevance":10},{"className":"attribute","begin":"^\\w","end":": ","excludeEnd":true,"starts":{"end":"$","relevance":0}},{"className":"literal","begin":"^-","end":"$"},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}`)
}