

<template>
  <div id="Editors">
    <Tabs>
      <Tab name="Hex Editor" :selected="true">
        <div>
          <HexEditor v-on:packet-modified="updateData($event)" :data="data"/>
        </div>
      </Tab>
      <Tab name="Text Editor">
        <TextEditor v-on:packet-modified="updateData($event)" :data="data"/> 
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
 /*
 data() {
    return { data: "\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC", hexVisible: true }
  },
*/
  methods: {
    updateData: function (data) {
      this.data = data
    }
  },
    computed: {
      internalData: {
      get() {
        return this.data;
      },
      set(data) {
        console.log("emitting packet-modified, data")
        console.log('emitting data from hex:', data)
        this.$emit('packet-modified', data)
      }
    }
  }
}

</script>

<style>

</style>
