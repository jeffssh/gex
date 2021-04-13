

<template>
  <div id="Editors">
    <Tabs>
      <Tab name="Hex Editor" :selected="true">
        <div>
          <HexEditor v-on:send-to-repeater="addRepeaterTab($event)" v-on:packet-modified="updateData($event)" :data="internalData"/>
        </div>
      </Tab>
      <Tab name="Text Editor">
        <TextEditor v-on:send-to-repeater="addRepeaterTab($event)" v-on:packet-modified="updateData($event)" :data="internalData"/> 
      </Tab>
    </Tabs>
  </div>
</template>

<script>
import TextEditor from './TextEditor.vue'
import HexEditor from './HexEditor.vue'
import Tab from './Tab.vue'
import Tabs from './Tabs.vue'

export default {
  name: 'Editors',
  components: {
    TextEditor,
    HexEditor,
    Tab,
    Tabs
  },
  props: {
    data: String
  },
  methods: {
    updateData: function (data) {
      // TODO need to emit here instead of modify
      this.$emit('packet-modified', data)
      //this.data = data
    },
    addRepeaterTab(data)  {
      this.$emit('send-to-repeater', data)
    }
  },
  computed: {
    internalData: {
      get() {
        //console.log(this.data.repeat(1))
        //console.log(this.data)
        return this.data.repeat(1)//Object.assign("", this.data)
      },
    }
  }
}

</script>

<style scoped>

</style>
