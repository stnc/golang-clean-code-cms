user da mail olayinda mail daha once kayitli mi ona balkilacak 


https://github.com/Pungyeon/clean-go-article#Test-Driven-Development


https://davicotico.github.io/jQuery-Menu-Editor/

https://www.bootdey.com/snippets/view/notification-list


# json notes -- htmlescape
https://stackoverflow.com/questions/57086035/go-template-htmlescape-json-data-and-it-show-34-issue
https://github.com/flosch/pongo2/issues/129
https://github.com/TylerBrock/colorjson

https://github.com/denisbakhtin/ginblog/blob/master/main.go burada memstore kullanımı var ona bakılablir

https://zetcode.com/golang/slice/
## aweseom lin 

https://github.com/dustin/go-humanize
https://github.com/russross/blackfriday

# go notes
## autoescape
bu kod pongo2 de json gibi data parse eder
ontrols the current auto-escaping behavior. This tag takes either on or off as an argument and that determines whether auto-escaping is in effect inside the block. The block is closed with an endautoescape ending tag.

When auto-escaping is in effect, all variable content has HTML escaping applied to it before placing the result into the output (but after any filters have been applied). This is equivalent to manually applying the escape filter to each variable.

The only exceptions are variables that are already marked as “safe” from escaping, either by the code that populated the variable, or because it has had the safe or escape filters applied.
  
                 {% autoescape off %}
                  {{json}} 
                {% endautoescape %};

## verbatim

Stops the template engine from rendering the contents of this block tag.

A common use is to allow a JavaScript template layer that collides with Django’s syntax. For example:

{% verbatim %}
{{if dying}}Still alive.{{/if}}
{% endverbatim %}


