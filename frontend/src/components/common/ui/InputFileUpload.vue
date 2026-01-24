<script setup lang="ts">
import { ref } from 'vue'

const props = withDefaults(
  defineProps<{
    label: string
    modelValue?: string | File | (File | string)[] | null
    multiple?: boolean
    accept?: string
    required?: boolean
    hasError?: boolean
  }>(),
  {
    multiple: false,
    modelValue: null,
    required: false,
    hasError: false,
  },
)

const emit = defineEmits<{
  'update:modelValue': [value: string | File | (File | string)[] | null]
}>()

const isDragging = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)

const handleFiles = (fileList: FileList) => {
  if (props.multiple) {
    const newFiles = Array.from(fileList)
    const currentFiles = Array.isArray(props.modelValue) ? props.modelValue : []
    emit('update:modelValue', [...currentFiles, ...newFiles])
  } else {
    emit('update:modelValue', fileList[0])
  }
}

const onDragOver = (e: DragEvent) => {
  e.preventDefault()
  isDragging.value = true
}

const onDragLeave = (e: DragEvent) => {
  e.preventDefault()
  isDragging.value = false
}

const onDrop = (e: DragEvent) => {
  e.preventDefault()
  isDragging.value = false
  const files = e.dataTransfer?.files
  if (files && files.length > 0) {
    handleFiles(files)
  }
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    handleFiles(target.files)
  }
}

const triggerBrowse = () => {
  fileInput.value?.click()
}

const removeFile = (index?: number) => {
  if (props.multiple && Array.isArray(props.modelValue) && typeof index === 'number') {
    const newFiles = [...props.modelValue]
    newFiles.splice(index, 1)
    emit('update:modelValue', newFiles)
  } else {
    emit('update:modelValue', null)
  }
  
  if (fileInput.value) fileInput.value.value = ''
}

const formatSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`
}
</script>

<template>
  <div class="field-group" :class="{ 'has-error': hasError }">
    <label class="field-label">
      {{ label }}
    </label>

    <div
      class="upload-container"
      :class="{ 'is-dragging': isDragging }"
      @dragover="onDragOver"
      @dragleave="onDragLeave"
      @drop="onDrop"
      @click="triggerBrowse"
    >
      <input
        ref="fileInput"
        type="file"
        class="file-input"
        :multiple="multiple"
        :accept="accept"
        @change="onFileChange"
      />

      <div class="upload-content">
        <div class="cloud-icon">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"></path>
            <path d="M12 12v9"></path>
            <path d="m16 16-4-4-4 4"></path>
          </svg>
        </div>
        <div class="upload-text">
          <span class="browse-text">Browse Files</span>
          <span class="drag-text">Drag and drop files here</span>
        </div>
      </div>
    </div>

    <div v-if="modelValue" class="file-list">
      <template v-if="Array.isArray(modelValue)">
        <div v-for="(file, index) in modelValue" :key="index" class="file-item">
          <div class="file-info">
            <span class="file-name">{{ typeof file === 'string' ? file : file.name }}</span>
            <span class="file-size" v-if="typeof file !== 'string'">{{
              formatSize(file.size)
            }}</span>
          </div>
          <button type="button" class="remove-btn" @click.stop="removeFile(index)">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M18 6 6 18"></path>
              <path d="m6 6 12 12"></path>
            </svg>
          </button>
        </div>
      </template>
      <template v-else>
        <div class="file-item">
          <div class="file-info">
            <span class="file-name">{{
              typeof modelValue === 'string' ? modelValue : (modelValue as File).name
            }}</span>
            <span class="file-size" v-if="typeof modelValue !== 'string'">{{
              formatSize((modelValue as File).size)
            }}</span>
          </div>
          <button type="button" class="remove-btn" @click.stop="removeFile()">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M18 6 6 18"></path>
              <path d="m6 6 12 12"></path>
            </svg>
          </button>
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped>
.field-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.field-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-primary);
}

.required {
  color: var(--color-danger);
  margin-left: 2px;
}

.upload-container {
  border: 1px dashed var(--border-color);
  border-radius: 8px;
  background-color: var(--text-inverse);
  box-shadow: 0 2px 4px rgb(0 0 0 / 10%);
  padding: 24px;
  text-align: center;
  cursor: pointer;
  transition: all 0.2s ease;
  min-height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-container:hover {
  background-color: hsl(from var(--color-neutral) h s 95%);
  border-color: hsl(from var(--color-neutral) h s 60%);
}

.upload-container.is-dragging {
  background-color: hsl(from var(--color-neutral) h s 90%);
  border-color: var(--color-primary);
}

.has-error .upload-container {
  border-color: var(--color-danger);
  background-color: hsl(from var(--color-danger) h s 95%);
}

.file-input {
  display: none;
}

.upload-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}

.cloud-icon {
  color: hsl(from var(--color-neutral) h s 60%);
  background: var(--border-color);
  padding: 8px;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-text {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.browse-text {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 1rem;
}

.drag-text {
  color: hsl(from var(--color-neutral) h s 40%);
  font-size: 0.875rem;
}

.file-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 4px;
}

.file-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  padding: 8px 12px;
}

.file-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.file-name {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-primary);
  word-break: break-all;
}

.file-size {
  font-size: 0.75rem;
  color: hsl(from var(--color-neutral) h s 40%);
}

.remove-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: hsl(from var(--color-neutral) h s 60%);
  padding: 4px;
  border-radius: 4px;
  transition: color 0.2s;
}

.remove-btn:hover {
  color: var(--color-danger);
  background-color: hsl(from var(--color-danger) h s 95%);
}
</style>
