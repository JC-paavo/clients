<template>
  <el-dialog
    class="base-dialog"
    :append-to-body="true"
    :modal-append-to-body="true"
    :title="title"
    :top="top"
    :width="width"
    v-bind="$attrs"
  >
    <slot/>
    <template #footer>
      <div class="dialog-footer">
        <el-button v-if="showCancel && showButtons" @click="onCancel">{{ cancelTitle }}</el-button>
        <el-button v-if="showSave && showButtons" type="success" @click="onSave">{{ saveTitle }}</el-button>
        <el-button v-if="showConfirm && showButtons" :loading="loadingStatus" type="primary" @click="onConfirm">
          {{ confirmTitle }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: 'SettingDialog',
  props: {
    title: {
      type: String,
      default: 'Title'
    },
    top: {
      type: String,
      default: '12vh'
    },
    width: {
      type: String,
      default: '70%'
    },
    showConfirm: {
      type: Boolean,
      default: true
    },
    confirmTitle: {
      type: String,
      default() {
        return "保存且默认"
      }
    },
    showCancel: {
      type: Boolean,
      default: true
    },
    cancelTitle: {
      type: String,
      default() {
        return "关闭"
      }
    },
    showSave: {
      type: Boolean,
      default: true
    },
    saveTitle: {
      type: String,
      default() {
        return "保存"
      }
    },
    showButtons: {
      type: Boolean,
      default: true
    },
    loadingStatus: {
      type: Boolean,
      default: false
    },
    maxWidth: {
      type: String,
      default: '1200px'
    }
  },
  data() {
    return {}
  },
  methods: {
    onCancel() {
      this.$emit('cancel')
    },
    onConfirm() {
      this.$emit('confirm')
    },
    onSave() {
      this.$emit('save')
    }
  }
}
</script>

<style lang="scss">
.base-dialog {
  .el-dialog__body {
    padding: 10px 20px;
  }
}
</style>

