<template>
  <div class="tabpane">
    
  </div>
</template>

<script>
export default {
  name: 'TabPane',
  props: {
    msg: String
  }
}

Vue.component('tabs', {
  template: `
    <div>
      <div class="tabs">
        <ul>
          <li v-for="tab in tabs" :class="{ 'is-active': tab.isActive }">
            <a :href="tab.href" @click="selectTab(tab)">{{ tab.name }}</a>
          </li>
        </ul>
      </div>

      <div class="tabs-details">
        <slot></slot>
      </div>
    </div>
  `,
  data() {
    return { tabs: [] }
  },
  created() {
    this.tabs = this.$children;
  },
  methods: {
    selectTab: function(selectedTab) {
      this.tabs.forEach(tab => {
        tab.isActive = (tab.name == selectedTab.name);
      });
    }
  }
});

Vue.component('tab', {
  props: {
    name: { required: true },
    selected: { default: false }
  },
  template: `<div v-show="isActive"><slot></slot></div>`,
  data() {
    return { isActive: false }
  },
  computed: {
    href() {
      return '#' + this.name.toLowerCase().replace(/ /g , '-');
    }
  },
  mounted() {
    this.isActive = this.selected;
  }
});

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
:root {
  padding-top: 40px;
}
</style>
