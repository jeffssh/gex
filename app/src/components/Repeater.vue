<template>
  <div class="Repeater">
   <div style="height: 50px;" ref="repeater">
      <buttons>
        <button id="dark-button" class="button is-dark" style="margin-left: 10px;" v-on:click="websocketSendTabData">Send</button>
      </buttons>
    </div>
    <ejs-tab id='dynamic_tab' ref='TabInstance' :animation="{previous: null, next: null}"  :selected="tabSelected">
      <e-tabitems>    
      </e-tabitems>
    </ejs-tab>
  <Editors v-on:send-to-repeater="addRepeaterTab($event)" v-on:packet-modified="updateData($event)" :data="tabData" ></Editors>
  </div>
</template>

<script>
import Tabs from './Tabs.vue'
import Tab from './Tab.vue'
import Vue from 'vue'
import Editors from './Editors.vue'


import { TabPlugin } from '@syncfusion/ej2-vue-navigations';
//import { createElement } from '@syncfusion/ej2-base';

Vue.use(TabPlugin);

export default {
  name: 'Repeater',
  components: {
    Editors,
  },
  props: {
    repeaterTabs: Object,
    ws: Object
  },
  methods: {
    tabSelected: function(e) {
      this.currTab = parseInt(e.selectedItem.innerText)
      this.tabData = this.repeaterTabs.get(this.currTab)
    },
    updateData: function(data) {
      this.repeaterTabs.set(this.currTab, data)
      this.tabData = data
    },
    websocketSendTabData: function() {
      var data = this.tabData
      var buf = new ArrayBuffer(data.length)
      var bufView = new Uint8Array(buf)
      for (var i=0, strLen=data.length; i < strLen; i++) {
        bufView[i] = data.charCodeAt(i)
      }
      this.ws.send(buf)
    },
  },
  mounted() {
  // https://ej2.syncfusion.com/vue/documentation/tab/how-to/load-tab-items-dynamically/
    var tabObj = this.$refs.TabInstance.ej2Instances;
    let totalItems = document.querySelectorAll('#dynamic_tab .e-toolbar-item').length;
    this.repeaterTabs.forEach(function(v, k) {
        let item =  { header: { text: k.toString() }, content: ""};
        tabObj.addTab([item], (totalItems++));
      }
    )
  },
  data() {
    return {
      tabData: "",
      currTab: 0,
      rerender: false
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  @import "../../node_modules/@syncfusion/ej2-base/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-buttons/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-popups/styles/material.css";
  @import "../../node_modules/@syncfusion/ej2-navigations/styles/material.css";

  #dynamic_tab {
    text-align: left;
  }
</style>
