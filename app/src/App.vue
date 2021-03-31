

<template>
  <div id="app">
    <button id="dark-button" class="button is-dark" style="margin-left: 10px;" :userObj="data" v-on:click="updateData('AAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDD')">Send Packet</button>
    <Tabs>
      <Tab name="TCP Intercept" :selected="true" ref="intercept">
        <div style="height: 50px;">
          <!-- https://bulma.io/documentation/elements/button/ --> 
          <buttons>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Forward</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Drop</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;">Intercept is off</button>
          </buttons>
        </div>
          
          <!--
            <Tab name="Hex Editor" :selected="true" v-on:click.native="clickHex()">
              <div>
                <HexEditor v-on:packet-modified="updateData($event)" :data="initdata"/>
              </div>
            </Tab>
            <Tab name="Text Editor">
              <TextEditor v-on:packet-modified="updateData($event)" :data="initdata"/> 
            </Tab>
            <Tab name="Settings">
            </Tab>
          -->
            <Editors :data=initdata v-on:send-to-repeater="addRepeaterTab($event)"> </Editors>
      </Tab>
      <Tab name="TCP Proxy"></Tab>
      <Tab name="TCP Repeater">
        <Repeater :repeaterTabs="repeaterTabs"></Repeater>
      </Tab>
    </Tabs>
  </div>
</template>

<script>
import Vue from 'vue'
import TextEditor from './components/TextEditor.vue'
import HexEditor from './components/HexEditor.vue'
import Editors from './components/Editors.vue'
//import Table from './components/Table.vue'

import Tab from './components/Tab.vue'
import Tabs from './components/Tabs.vue'
import Repeater from './components/Repeater.vue'

export default {
  name: 'App',
  components: {
    Editors,
    Tab,
    Tabs,
    Repeater
  },
    data() {
    return { 
      initdata: "AAAA",//"\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC\xCC",
      repeaterTabs: new Map(),
      numTabs: 0
    }
  },
  methods: {
    addRepeaterTab(data)  {
      this.repeaterTabs.set(++this.numTabs, data)      
      /*
      let TabClass = Vue.extend(Tab)
      let e = new TabClass(
        {
            propsData: { 
              name: "test",//this.repeaterTabs.length,
              selected: true,
              data: data 
            }
        }
      )
      e.$slots.default = [ 'Click me!' ]
      this.repeaterTabs.push(e)
      */
    }
  }
}

</script>

<style>
@import "css/bulma.min.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  overflow: auto;
}

#dark-button {
  float: left;
  padding: 10px
}

</style>
