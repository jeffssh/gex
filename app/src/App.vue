

<template>
  <div id="app">
    <!--
    <button id="ws-test" class="button is-dark" style="margin-left: 10px;"  v-on:click="websocketSend(interData)">ws test</button>
    <button id="dark-button" class="button is-dark" style="margin-left: 10px;" :userObj="data" v-on:click="updateData('AAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDDAAAABBBBCCCCDDDD')">Send Packet</button>
    -->
    <Tabs>
      <Tab name="TCP Intercept" :selected="true" ref="intercept">
        <div style="height: 50px;">
          <!-- https://bulma.io/documentation/elements/button/ --> 
          <buttons>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;" v-on:click="forwardPacket()">Forward</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;" v-on:click="dropPacket()">Drop</button>
            <button id="dark-button" class="button is-dark" style="margin-left: 10px;" v-on:click="toggleIntercept()">Intercept is {{intercept  ? "on" : "off"}} | Queue: {{packetQueue.length}}</button>
          </buttons>
        </div>
        <Editors v-on:packet-modified="updateData($event)" :data=interData v-on:send-to-repeater="addRepeaterTab($event)"> </Editors>
      </Tab>
      <!--<Tab name="TCP Proxy"></Tab>-->
      <Tab name="TCP Repeater">
        <Repeater :repeaterTabs="repeaterTabs" :ws=ws></Repeater>
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
      interData: "",
      repeaterTabs: new Map(),
      numTabs: 0,
      ws: null,
      intercept: false,
      shownOnePacket: false,
      packetQueue: [],
    }
  },
  created () {
    document.title = "Gex"
    var ws_url = "ws://" + window.location.href.split("/")[2] + "/ws"
    this.ws = new WebSocket(ws_url)
    this.ws.binaryType = "arraybuffer"
    this.ws.onmessage = this.websocketRecv
  },
  methods: {
    addRepeaterTab(data)  {
      this.repeaterTabs.set(++this.numTabs, data)      
    },
    updateData: function(data) {
      this.interData = data == undefined ? "" : data
    },
    getNextInQueue: function() {
      this.updateData(this.packetQueue.pop())
    },
    dropPacket: function() {
      if(!this.packetQueue.length) {
        this.shownOnePacket = false
      }
      this.getNextInQueue()
    },
    forwardPacket: function() {
      if(this.interData != "") {
        this.websocketSend(this.interData);
      }
      if(!this.packetQueue.length) {
        this.shownOnePacket = false
      }
      this.getNextInQueue()
    },
    toggleIntercept: function() {
      this.intercept = !this.intercept
      if(this.intercept) { 
        this.shownOnePacket = false
        this.getNextInQueue() 
      } else {
        if(this.interData != "") {
          console.log("sending current intercept data")
          this.websocketSend(this.interData);
        }
        while(this.packetQueue.length){
          this.websocketSend(this.packetQueue.pop())
        }
        this.updateData(undefined)
      }
    },
    websocketSend: function(data) {
      var buf = new ArrayBuffer(data.length)
      var bufView = new Uint8Array(buf)
      for (var i=0, strLen=data.length; i < strLen; i++) {
        bufView[i] = data.charCodeAt(i)
      }
      this.ws.send(buf)
    },
    websocketRecv (event) {
        var s = String.fromCharCode.apply(null, new Uint8Array(event.data))
        this.packetQueue.unshift(s)
        if(this.intercept && !this.shownOnePacket) { 
          this.shownOnePacket = true
          this.getNextInQueue()
        }
        if(!this.intercept) {
          this.websocketSend(this.packetQueue.pop())
        }
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
  width: 100%;
  overflow: hidden;
}

#dark-button {
  float: left;
  padding: 10px
}
</style>
