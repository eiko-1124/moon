<template>
  <div class="edit-mood">
    <section>
      <div class="edit-mood-box">
        <h3>我的心情</h3>
        <el-button
          class="edit-mood-submit"
          type="primary"
          @click="submit"
          :disabled="updateFlag"
          >发布</el-button
        >
      </div>
      <div class="edit-mood-box">
        <el-input
          v-model="moodText"
          type="textarea"
          :autosize="{ minRows: 3, maxRows: 5 }"
        ></el-input>
      </div>
      <div class="edit-mood-box edit-update">
        <el-scrollbar style="background-color: white" v-loading="updateFlag">
          <el-upload
            v-model:file-list="fileList"
            list-type="picture-card"
            :limit="9"
            :auto-upload="false"
            :on-change="beforeAvatarUpload"
            :on-preview="handlePictureCardPreview"
            ref="elUpload"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>

          <el-dialog v-model="dialogVisible">
            <img
              style="width: 100%"
              :src="dialogImageUrl"
              alt="Preview Image"
            />
          </el-dialog>
        </el-scrollbar>
      </div>
    </section>
    <section>
      <mood-list></mood-list>
    </section>
  </div>
</template>

<script>
import { reactive, ref } from "@vue/reactivity";
import { moodPicture } from "@/api/apis/file";
import { MoodSubmit } from "@/api/apis/mood";
import MoodList from "@/components/list/MoodList.vue";
import { useStore } from "vuex";
export default {
  components: { MoodList },
  setup() {
    const store = useStore();
    const dialogVisible = ref(false);
    const dialogImageUrl = ref("");
    const fileList = reactive([]);
    const moodText = ref("");
    const elUpload = ref(null);
    const updateFlag = ref(false);
    const beforeAvatarUpload = (file) => {
      if (!/image/.test(file.raw.type)) {
        ElMessage.error("Avatar picture must be image!");
        elUpload.value.handleRemove(file);
        return;
      }
      if (file.raw.size > 1024 * 1024 * 15) {
        ElMessage.error("Avatar picture size must down to 15m!");
        elUpload.value.handleRemove(file);
      }
    };
    const handlePictureCardPreview = (uploadFile) => {
      dialogVisible.value = true;
      dialogImageUrl.value = uploadFile.url;
    };
    const submit = () => {
      if (moodText.value == "") return;
      updateFlag.value = true;
      let picArr = [];
      if (fileList.length != 0) {
        picUpdate(picArr, fileList, 0, fileList.length)
          .then(() => {
            let picStr = "";
            for (let id of picArr) {
              picStr += id;
              picStr += "#";
            }
            const form = new FormData();
            form.append("text", moodText.value);
            form.append("pic", picStr.substring(0, picStr.length - 1));
            textUpdate(form);
          })
          .catch(() => {
            updateFlag.value = false;
          });
      } else {
        const form = new FormData();
        form.append("text", moodText.value);
        form.append("pic", "");
        textUpdate(form);
      }
    };
    const picUpdate = (pArr, files, count, sum) => {
      return new Promise((resolve, reject) => {
        if (count == sum) resolve();
        const form = new FormData();
        form.append("file", files[count].raw);
        moodPicture(
          form,
          (Res) => {
            if (Res.res == 1) {
              pArr.push(Res.id);
              picUpdate(pArr, files, count + 1, sum).then(() => resolve());
            } else reject();
          },
          (Err) => {
            console.log(Err);
            reject();
          }
        );
      });
    };
    const textUpdate = (form) => {
      MoodSubmit(
        form,
        (Res) => {
          if (Res.res == 1) {
            updateFlag.value = false;
            store.dispatch("submitMood");
            ElMessage("发布成功");
            moodText.value = "";
            elUpload.value.clearFiles();
          } else ElMessage("发布失败");
        },
        (Err) => {
          console.log(Err);
          ElMessage("发布失败");
        }
      );
    };
    return {
      beforeAvatarUpload,
      handlePictureCardPreview,
      submit,
      elUpload,
      dialogVisible,
      dialogImageUrl,
      fileList,
      moodText,
      updateFlag,
    };
  },
};
</script>

<style lang="scss" scoped>
.edit-mood {
  display: flex;
  min-width: 23.5rem;
  min-height: calc(100vh - 9.5rem);
  & > section {
    transition-duration: 0.5s;
    animation: load-bottom 1s;
  }
  & > section:nth-child(2) {
    flex: 2;
    height: calc(100vh - 9.5rem);
  }
  & > section:nth-child(1) {
    flex: 3;
    padding: 1rem;
    margin-right: 1rem;
    height: calc(100vh - 9.5rem);
    background-color: white;
    box-shadow: 0px 0px 1px 0px darkgray;
  }
}
.edit-mood-box {
  position: relative;
  margin-bottom: 1rem;
  h3 {
    color: darkcyan;
  }
}
.edit-mood-submit {
  position: absolute;
  top: 0rem;
  right: 0rem;
}
.edit-update {
  height: calc(100% - 8rem);
}
@media screen and (max-width: 960px) {
  .edit-mood {
    min-width: 40rem;
    & > section:nth-child(2) {
      flex: 3;
    }
  }
}
@media screen and (max-width: 560px) {
  .edit-mood {
    min-width: 23.5rem;
    flex-direction: column;
    & > section:nth-child(1) {
      margin-bottom: 1rem;
    }
  }
}
</style>