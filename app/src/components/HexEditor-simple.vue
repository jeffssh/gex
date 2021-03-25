<template>
<div id="hexeditor" class="hexeditor">
<body onload='
// Reset the textarea value
m.value="00010203040506008090a0b0c0d0e0f";
// Init the top cell content
for(i=0;i<16;i++)
  t.innerHTML+=(0+i.toString(16)).slice(-2)+" ";
'>
</body>
<table border>
  <tr>
    <td>
    <pre></pre>
  </td>
    <td id=t></td>
  </tr>
  <tr>
    <td id=l width=80>00000000</td>
  <td>
  <textarea spellcheck=false id=m oninput='
  // On input, store the length of clean hex before the textarea caret in b
  b=value
  .substr(0,selectionStart)
  .replace(/[^0-9A-F]/ig,"")
  .replace(/(..)/g,"$1 ")
  .length;
  // Clean the textarea value
  value=value
  .replace(/[^0-9A-F]/ig,"")
  .replace(/(..)/g,"$1 ")
  .replace(/ $/,"")
  .toUpperCase();
  // Set the height of the textarea according to its length
  style.height=(1.5+value.length/47)+"em";
  // Reset h
  h="";
  // Loop on textarea lines
  for(i=0;i<value.length/48;i++)
    
    // Add line number to h
    h+=(1E7+(16*i).toString(16)).slice(-8)+" ";
  // Write h on the left column
  l.innerHTML=h;
  // Reset h
  h="";
  // Loop on the hex values
  for(i=0;i<value.length;i+=3)
    
    // Convert them in numbers
    c=parseInt(value.substr(i,2),16),
    
    // Convert in chars (if the charCode is in [64-126] (maybe more later)) or ".".
    h=63<c&&127>c?h+String.fromCharCode(c):h+".";
    
  // Write h in the right column (with line breaks every 16 chars)
  r.innerHTML=h.replace(/(.{16})/g,"$1 ");
  // If the caret position is after a space or a line break, place it at the previous index so we can use backspace to erase hex code
  if(value[b]==" ")
    b--;
  // Put the textarea caret at the right place
  setSelectionRange(b,b)'
  cols=48></textarea>
  </td>
  <td width=160 id=r>.</td>
  </tr>
</table>
</div>
</template>

<script>
export default {
  methods: {
    updateData(data) {
      var buf = new ArrayBuffer(data.length); // 2 bytes for each char
      var bufView = new Uint8Array(buf);
      for (var i=0, strLen=data.length; i < strLen; i++) {
        bufView[i] = data.charCodeAt(i);
      }
      this.dataView = new DataView(buf);
      //this.updateInterpreter();
    },
  },
  name: 'HexEditor',
  props: {
    data: String
  },
  watch: {
    data: function (val) {
      this.updateData(val)
    },
  },
  mounted() {
    this.updateData(this.data)
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
*{margin:0;padding:0;vertical-align:top;font:1em/1em courier}
#m{height:1.5em;resize:none;overflow:hidden}
#t{padding:0 2px}
#w{position:absolute;opacity:.001}
</style>
