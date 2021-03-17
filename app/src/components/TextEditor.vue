<template>
  <div class="texteditor">
    <AceEditor 
        v-model="content" 
        @init="editorInit" 
        lang="text" 
        theme="monokai" 
        width="100%" 
        height="600px"
        :options="{
            fontSize: 14,
            highlightActiveLine: true,
            showLineNumbers: true,
            tabSize: 4,
            showGutter: true,
        }"
        :commands="[
            {
                name: 'save',
                bindKey: { win: 'Ctrl-s', mac: 'Command-s' },
                exec: dataSubmit,
                readOnly: true,
            },
        ]"
    />

  </div>
</template>

<script>
import AceEditor from 'vuejs-ace-editor';

export default {
  name: 'TextEditor',
   components: {
    AceEditor
  },
   methods: {
        dataSubmit() {
            alert("submit")
        },
        editorInit: function () {
            require('brace/mode/html')                
            require('brace/theme/monokai')
        }
    }/*,
  props: {
    data: String
  }*/
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#editor {
  width: 600vh;
  height: 50vh;
 }
</style>
