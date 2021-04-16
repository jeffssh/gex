<template>
  <div id="hexeditor" class="hexeditor">
  <div class="drag" v-if="!dataView">
    <div class="open-file">
      <input class="open-file-input" id="file" type="file" name="file" @change="handleOpenFile"/>
      <label class="open-file-label" for="file">Drag and drop a file into this area or click here</label>
    </div>
  </div>
  <div class="file" v-if="dataView">
    <div class="offsets" title="Offset">
      <div class="offset" v-for="(offset, lineIndex) of offsets" :key="lineIndex" :class="{ 'active': isLineActive(lineIndex) }" @click="handleLineClick(lineIndex, $event)">{{offset}}</div>
    </div>
    <div class="lines hex" title="Hex value">
      <div class="line" v-for="(line, lineIndex) of hex" :key="lineIndex" :class="{ 'active': isLineActive(lineIndex) }" @click="handleLineClick(lineIndex, $event)">
        <div class="value" v-for="(value, valueIndex) of line" :key="(lineIndex * 16) + valueIndex" :class="{ 'active': isValueActive(lineIndex, valueIndex) }" @click="handleValueClick(valueIndex, $event)">{{value}}</div>
      </div>
    </div>
    <div class="lines ascii" title="Ascii value">
      <div class="line" v-for="(line, lineIndex) of ascii" :key="lineIndex" :class="{ 'active': isLineActive(lineIndex) }" @click="handleLineClick(lineIndex, $event)">
        <div class="value" v-for="(value, valueIndex) of line" :key="(lineIndex * 16) + valueIndex" :class="{ 'active': isValueActive(lineIndex, valueIndex) }" @click="handleValueClick(valueIndex, $event)">{{value}}</div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
//console.log(this.$parent)
//console.log(this.$parent.selected)
export default {
  created() {
    document.addEventListener('drag', this.handleDrag);
    document.addEventListener('dragstart', this.handleDrag);
    document.addEventListener('dragenter', this.handleDrag);
    document.addEventListener('dragexit', this.handleDrag);
    document.addEventListener('dragend', this.handleDrag);
    document.addEventListener('dragleave', this.handleDrag);
    document.addEventListener('dragover', this.handleDrag);
    document.addEventListener('drop', this.handleDrag);
    document.addEventListener('keydown', this.handleKey);
    document.addEventListener('wheel', this.handleWheel);
    window.addEventListener('resize', this.handleResize);
    window.addEventListener('keypress', this.handleKeyPress);
  },
  destroyed() {
    window.removeEventListener('keypress', this.handleKeyPress);
  },
  methods: {
    handleKeyPress(e) {
      // handle keys
      let inKey = String.fromCharCode(e.keyCode).toUpperCase();
      switch(inKey) {
      case "R": {
        // r (repeater)
        this.$emit('send-to-repeater', this.internalData)
        break;
      }
      default: {
        // code block
        // handle byte insertion
        let alphabet = "0123456789ABCDEF"
        if (!alphabet.includes(inKey)) {return}
        this.input += inKey;
        if (this.input.length == 2) {  
          let index = this.column + this.row.current * 0x10
          let ch = String.fromCharCode(parseInt(this.input, 16))
          this.internalData =  this.internalData.substring(0,index) + ch +  this.internalData.substring(index+1);
          //console.log(this.internalData.substring(0,index) + ch +  this.internalData.substring(index+1))
          this.input = ''
        }
      }
    }


    },
    updateData(data) {
      var buf = new ArrayBuffer(data.length); // 2 bytes for each char
      var bufView = new Uint8Array(buf);
      for (var i=0, strLen=data.length; i < strLen; i++) {
        bufView[i] = data.charCodeAt(i);
      }
      this.dataView = new DataView(buf);
      this.updateInterpreter();
    },
    getRows(fn) {
      if (!this.dataView) {
        return [];
      }

      const rows = [];

      for (let row = this.row.start; row < this.row.start + this.rows; row++) {
        const values = [];

        for (let column = 0; column < this.rowLength; column++) {
          const offset = row * this.rowLength + column;
          const value = offset < this.dataView.byteLength ? this.dataView.getUint8(offset) : '  ';
          values.push(fn(value));
        }

        rows.push(values);
      }

      return rows;
    },

    updateColumnToLine() {
      const offset = this.row.current * this.rowLength + this.column;

      if (offset >= this.dataView.byteLength) {
        this.column = this.dataView.byteLength % this.rowLength - 1;
      }
    },

    moveCharLeft() {
      this.column = Math.max(0, this.column - 1);
      this.updateInterpreter();
    },

    moveCharRight() {
      const newColumn = Math.min(this.rowLength - 1, this.column + 1);
      const offset = this.row.current * this.rowLength + newColumn;

      if (offset < this.dataView.byteLength) {
        this.column = newColumn;
      }

      this.updateInterpreter();
    },

    moveLineUp() {
      this.row.current = Math.max(0, this.row.current - 1);

      if (this.row.current < this.row.start) {
        this.row.start = this.row.current;
      }

      this.updateColumnToLine();
      this.updateInterpreter();
    },

    moveLineDown() {
      this.row.current = Math.min(this.maxRows, this.row.current + 1);

      if (this.row.current > this.row.start + (this.rows - 1)) {
        this.row.start = this.row.current - (this.rows - 1);
      }

      this.updateColumnToLine();
      this.updateInterpreter();
    },

    movePageUp() {
      this.row.start = Math.max(0, this.row.start - this.rows);
      this.row.current = this.row.start;
      this.updateColumnToLine();
      this.updateInterpreter();
    },

    movePageDown() {
      this.row.start = Math.min(this.maxStartRow, this.row.start + this.rows);
      this.row.current = this.row.start;
      this.updateColumnToLine();
      this.updateInterpreter();
    },

    moveToStart() {
      this.row.current = this.row.start = 0;
    },

    moveToEnd() {
      this.row.start = this.maxStartRow;
      this.row.current = this.maxRows;
    },

    goToChar(charIndex) {
      this.column = charIndex;
    },

    goToLineRelative(lineIndex) {
      this.row.current = this.row.start + lineIndex;
    },

    isValueActive(lineIndex, valueIndex) {
      if (!this.isLineActive(lineIndex)) {
        return false;
      }

      return valueIndex === this.column;
    },

    isLineActive(lineIndex) {
      return lineIndex === this.row.current - this.row.start;
    },

    updateInterpreter() {
      this.interpreter.u8 = this.dataView.getUint8(this.offset);
      this.interpreter.i8 = this.dataView.getInt8(this.offset);
      this.interpreter.u16le = this.dataView.getUint16(this.offset, true);
      this.interpreter.i16le = this.dataView.getInt16(this.offset, true);
      this.interpreter.u16be = this.dataView.getUint16(this.offset, false);
      this.interpreter.i16be = this.dataView.getInt16(this.offset, false);
      this.interpreter.u32le = this.dataView.getUint32(this.offset, true);
      this.interpreter.i32le = this.dataView.getInt32(this.offset, true);
      this.interpreter.u32be = this.dataView.getUint32(this.offset, false);
      this.interpreter.i32be = this.dataView.getInt32(this.offset, false);
      this.interpreter.f32le = this.dataView.getFloat32(this.offset, true);
      this.interpreter.f32be = this.dataView.getFloat32(this.offset, true);
      this.interpreter.f64le = this.dataView.getFloat64(this.offset, true);
      this.interpreter.f64be = this.dataView.getFloat64(this.offset, true);
      this.interpreter.tu32le = new Date(this.dataView.getUint32(this.offset, true));
      this.interpreter.tu32be = new Date(this.dataView.getUint32(this.offset, false));
      this.interpreter.tf32le = new Date(this.dataView.getFloat32(this.offset, true));
      this.interpreter.tf32be = new Date(this.dataView.getFloat32(this.offset, false));
      this.interpreter.tf64le = new Date(this.dataView.getFloat64(this.offset, true));
      this.interpreter.tf64be = new Date(this.dataView.getFloat64(this.offset, false));
    },

    loadFile(file) {
      const fr = new FileReader();
      fr.addEventListener('load', this.handleFile);
      fr.addEventListener('error', this.handleFile);
      fr.readAsArrayBuffer(file);
    },

    handlePalette(e) {
      const files = e.target.files;
      const [file] = files; //this.loadPalette(file)
    },

    handleOpenFile(e) {
      const files = e.target.files;
      const [file] = files;
      this.loadFile(file);
    },

    handleValueClick(valueIndex, e) {
      this.goToChar(valueIndex);
    },

    handleLineClick(lineIndex, e) {
      this.goToLineRelative(lineIndex);
    },

    handleResize(e) {
      //console.log(e);
    },

    handleKey(e) {
      if (e.code === 'ArrowUp' || e.key === 'k') {
        this.moveLineUp();
      } else if (e.code === 'ArrowDown' || e.key === 'j') {
        this.moveLineDown();
      }

      if (e.code === 'ArrowLeft' || e.key === 'h') {
        this.moveCharLeft();
      } else if (e.code === 'ArrowRight' || e.key === 'l') {
        this.moveCharRight();
      }

      if (e.code === 'PageUp') {
        this.movePageUp();
      } else if (e.code === 'PageDown') {
        this.movePageDown();
      }

      if (e.key === 'g') {
        this.moveToStart();
      } else if (e.key === 'G') {
        this.moveToEnd();
      }
    },

    handleWheel(e) {
      if (e.deltaY < 0) {
        this.moveLineUp();
      } else {
        this.moveLineDown();
      }
    },

    handleFile(e) {
      const fr = e.target;
      const arrayBuffer = fr.result;
      fr.removeEventListener('load', this.handleFile);
      fr.removeEventListener('error', this.handleFile);
      this.dataView = new DataView(arrayBuffer);
      this.updateInterpreter();
    },

    handleDrag(e) {
      if (e.type === 'dragover') {
        e.preventDefault();
      } else if (e.type === 'drop') {
        e.preventDefault();
        this.row.start = 0;
        this.row.current = 0;
        this.column = 0;
        const [file] = e.dataTransfer.files;
        this.loadFile(file);
      }
    }

  },

  data() {
    return {
      dataView: null,
      rowLength: 16,
      row: {
        start: 0,
        current: 0
      },
      column: 0,
      settings: {
        le: true,
        be: true,
        u: true,
        i: true,
        colorize: {
          enabled: false,
          palette: null
        }
      },
      interpreter: {
        u8: 0,
        i8: 0,
        u16le: 0,
        i16le: 0,
        u16be: 0,
        i16be: 0,
        u32le: 0,
        i32le: 0,
        u32be: 0,
        i32be: 0,
        f32le: 0,
        f32be: 0,
        f64le: 0,
        f64be: 0,
        tu32le: 0,
        tu32be: 0,
        tf32le: 0,
        tf32be: 0,
        tf64le: 0,
        tf64be: 0,
        binary: 0,
        hex: 0
      },
      input: ''
    };
  },

  computed: {
    internalData: {
      get() {
        return this.data;
      },
      set(data) {
        this.$emit('packet-modified', data)
      }
    },
    size() {
      if (!this.dataView) {
        return 0;
      }

      return this.dataView.byteLength;
    },

    offset() {
      return this.row.current * this.rowLength + this.column;
    },

    maxStartRow() {
      return this.maxRows - (this.rows - 1);
    },

    maxRows() {
      if (!this.dataView) {
        return 0;
      }

      return Math.floor(this.dataView.byteLength / this.rowLength);
    },

    rows() {
      return Math.ceil(this.$el.clientHeight / 16); // NOTE: This is not the row length (it's the pixel height of each line)
    },

    offsets() {
      if (!this.dataView) {
        return [];
      }

      const rows = [];

      for (let row = this.row.start; row < this.row.start + this.rows; row++) {
        rows.push((row * this.rowLength).toString(16).padStart(8, '0'));
      }

      return rows;
    },

    hex() {
      return this.getRows(value => value.toString(16).padStart(2, '0').toUpperCase());
    },

    ascii() {
      return this.getRows(function(value) {
          if (value == 0x20) { 
            //NARROW NO-BREAK SPACE hack
            return  "\u202F" 
          } 
          else {
            return value >= 32 && value <= 127 ? String.fromCharCode(value) : '.'
          }
        }
      );
    }

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
  },
}

DataView.prototype.getUint64 = function (byteOffset, littleEndian) {
  // split 64-bit number into two 32-bit (4-byte) parts
  const left = this.getUint32(byteOffset, littleEndian);
  const right = this.getUint32(byteOffset + 4, littleEndian);

  // combine the two 32-bit values
  const combined = littleEndian ? left + 2 ** 32 * right : 2 ** 32 * left + right;

  if (!Number.isSafeInteger(combined)) {
    console.warn(combined, 'exceeds MAX_SAFE_INTEGER. Precision may be lost');
  }

  return combined;
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
html,
body,
#hexeditor {
  width: 100%;
  height: 100%;
  display: flex;
  font-family: monospace;
  font-size: 1rem;
}
#hexeditor {
  overflow: hidden;
}
input[type="file"] {
  width: 0;
  height: 0;
  overflow: hidden;
  opacity: 0;
}
.button {
  background: #bbb;
  color: #333;
  border-radius: 0.5rem;
  padding: 1rem;
}
#hexeditor {
  flex-direction: row;
}
.file {
  display: flex;
  flex-direction: row;
}
.drag {
  background: #333;
  color: #bbb;
  display: flex;
  align-items: center;
  justify-content: center;
}
.drag .open-file {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 16rem 2rem;
  border: 2px dashed #bbb;
  border-radius: 1rem;
}
.drag .open-file label.open-file-label {
  font-size: 2rem;
  width: 100%;
  height: 100%;
  text-align: center;
}
.file,
.drag,
.offsets,
.values,
.interpreter {
  width: 100%;
  height: 100%;
}
.offsets {
  background: #222;
  color: #aaa;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  width: auto;
  height: auto;
}
.offsets .offset {
  margin-bottom: 0.125rem;
  padding: 0 1rem;
}
.offsets .offset.active {
  background: #224;
  color: #bbf;
}
.lines.hex {
  background: #333;
  color: #bbb;
  display: flex;
  flex-direction: column;
}
.lines.hex .line {
  margin-bottom: 0.125rem;
  display: flex;
  flex-direction: row;
  padding: 0 1rem;
}
.lines.hex .line.active {
  background: #336;
  color: #bbf;
}
.lines.hex .line .value {
  margin: 0 0.5rem;
}
.lines.hex .line .value.active {
  background: #bbf;
  color: #336;
}
.lines.ascii {
  background: #373737;
  color: #bbb;
  display: flex;
  flex-direction: column;
}
.lines.ascii .line {
  margin-bottom: 0.125rem;
  display: flex;
  flex-direction: row;
  padding: 0 1rem;
}
.lines.ascii .line.active {
  background: #373766;
  color: #bbf;
}
.lines.ascii .line .value.active {
  background: #fbf;
  color: #373766;
}
.interpreter {
  overflow-y: auto;
  padding-right: 1rem;
  padding-left: 1rem;
  background: #444;
  color: #bbb;
}
.interpreter form {
  width: 100%;
  margin-bottom: 1rem;
}
.interpreter form .settings-row {
  display: flex;
  flex-direction: row;
}
.interpreter form .setting {
  display: flex;
  flex: 1 1 auto;
  max-width: 12rem;
}
.interpreter form legend {
  padding: 0 1rem;
}
.interpreter form fieldset {
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;
}
.interpreter form .group {
  display: flex;
}
.interpreter table {
  width: 100%;
  border-collapse: collapse;
}
.interpreter table tr:nth-child(2n+0) td {
  padding-top: 1rem;
}
.interpreter table tr:nth-child(2n+1) td {
  padding-bottom: 1rem;
  border-bottom: 1px solid #bbb;
}
.interpreter table tr th {
  border-bottom: 1px solid #bbb;
  padding-bottom: 0.5rem;
}
.interpreter table tr td,
.interpreter table tr th {
  text-align: left;
  padding-right: 1rem;
}
.interpreter table tr td:last-child,
.interpreter table tr th:last-child {
  padding-right: 0;
}
.interpreter table tr td:first-child,
.interpreter table tr th:first-child {
  text-align: right;
}
.interpreter .u8 .type,
.interpreter .i8 .type {
  color: #c66;
}
.interpreter .u8 .value,
.interpreter .i8 .value {
  color: #df9f9f;
}
.interpreter .u16le .type,
.interpreter .i16le .type {
  color: #c96;
}
.interpreter .u16le .value,
.interpreter .i16le .value {
  color: #dfbf9f;
}
.interpreter .u16be .type,
.interpreter .i16be .type {
  color: #cc6;
}
.interpreter .u16be .value,
.interpreter .i16be .value {
  color: #dfdf9f;
}
.interpreter .u32le .type,
.interpreter .i32le .type {
  color: #9c6;
}
.interpreter .u32le .value,
.interpreter .i32le .value {
  color: #bfdf9f;
}
.interpreter .u32be .type,
.interpreter .i32be .type {
  color: #6c6;
}
.interpreter .u32be .value,
.interpreter .i32be .value {
  color: #9fdf9f;
}
.interpreter .f32le .type,
.interpreter .f32be .type {
  color: #6c9;
}
.interpreter .f32le .value,
.interpreter .f32be .value {
  color: #9fdfbf;
}
.interpreter .f64le .type,
.interpreter .f64be .type {
  color: #6cc;
}
.interpreter .f64le .value,
.interpreter .f64be .value {
  color: #9fdfdf;
}
.interpreter .tu32le .type,
.interpreter .tu32be .type {
  color: #69c;
}
.interpreter .tu32le .value,
.interpreter .tu32be .value {
  color: #9fbfdf;
}
.interpreter .tf32le .type,
.interpreter .tf32be .type {
  color: #7070c2;
}
.interpreter .tf32le .value,
.interpreter .tf32be .value {
  color: #a6a6d9;
}
.interpreter .tf64le .type,
.interpreter .tf64be .type {
  color: #9970c2;
}
.interpreter .tf64le .value,
.interpreter .tf64be .value {
  color: #bfa6d9;
}

</style>
