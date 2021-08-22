<template>
  <a-form-model :model="form" labelAlign="left" :label-col="labelCol" :wrapper-col="wrapperCol">
    <a-form-model-item label="卡片类型">
      <a-radio-group
        v-model="form.type"
        default-value="练习情况"
        button-style="solid"
      >
        <a-radio-button value="practice"> 练习情况 </a-radio-button>
        <a-radio-button value="guzhi"> 咕值信息 </a-radio-button>
      </a-radio-group>
    </a-form-model-item>

    <a-form-model-item label="用户 UID">
      <a-input type="number" v-model.number="form.uid" />
    </a-form-model-item>

    <a-form-model-item
      v-show="form.type == 'guzhi'"
      :label="itm"
      v-for="(itm, idx) in guzhiItems"
      :key="'guzhi' + itm"
    >
      <a-row>
        <a-col :span="20">
          <a-slider v-model="form.guzhi[idx]" :min="0" :max="100" />
        </a-col>
        <a-col :span="4">
          <a-input-number
            v-model="form.guzhi[idx]"
            :min="0"
            :max="100"
            style="marginleft: 16px"
          />
        </a-col>
      </a-row>
    </a-form-model-item>

    <a-form-model-item label="暗黑模式">
      <a-switch v-model="form.darkMode" />
    </a-form-model-item>

    <a-form-model-item label="隐藏标题">
      <a-switch v-model="form.hideTitle" />
    </a-form-model-item>

    <a-form-model-item label="卡片宽度">
      <a-row>
        <a-col :span="20">
          <a-slider v-model="form.cardWidth" :min="500" :max="1920" />
        </a-col>
        <a-col :span="4">
          <a-input-number
            v-model="form.cardWidth"
            :min="500"
            :max="1920"
            style="marginleft: 16px"
          />
        </a-col>
      </a-row>
    </a-form-model-item>

    <a-form-model-item label="效果预览">
      <img alt="" :src="imgUrl" />
    </a-form-model-item>

    <a-form-model-item label="复制代码">
      <a-tabs v-model="codeMode" @change="copyCode">
        <a-tab-pane
          v-for="codeMode in codeModes"
          :key="codeMode"
          :tab="codeMode"
        >
          <pre><code>{{codes[codeMode]}}</code></pre>
        </a-tab-pane>
      </a-tabs>
    </a-form-model-item>
  </a-form-model>
</template>

<script>
import { debounce } from "debounce";
import copy from 'copy-to-clipboard';

export default {
  data() {
    return {
      labelCol: { span: 3 },
      wrapperCol: { span: 21},
      guzhiItems: ["基础信用", "练习情况", "社区贡献", "比赛情况", "获得成就"],
      form: {
        type: "practice",
        uid: 313209,
        guzhi: [0, 0, 0, 0, 0],
        darkMode: false,
        hideTitle: false,
        cardWidth: 500,
      },
      codeModes: ["Markdown", "HTML", "URL"],
      codeMode: "Markdown",
      codes: {
        "Markdown": '',
        "HTML": '',
        "URL": '',
      },
      imgUrl: "https://luogu.wao3.cn/api/practice?id=313209",
    };
  },
  methods: {
    copyCode() {
      const code = this.codes[this.codeMode];
      copy(code);
      this.$message.success('已复制到剪切板');
    }
  },
  watch: {
    form: {
      handler() {
        const url = this.realtimeImgUrl;
        debounce(() => {
          this.imgUrl = url;
        }, 200)();
        this.codes["Markdown"] = `![我的练习情况](${url})`;
        this.codes["HTML"] = `<img src="${url}" alt="我的练习情况"/>`;
        this.codes["URL"] = url;
      },
      deep: true,
      immediate: true,
    }
  },
  computed: {
    realtimeImgUrl() {
      const form = this.form;
      let url = `https://luogu.wao3.cn/api/${form.type}?id=${form.uid}`;
      if (form.type == 'guzhi') {
        url += `&scores=${form.guzhi.join(',')}`;
      }
      if (form.darkMode) {
        url += '&dark_mode=true';
      }
      if (form.hideTitle) {
        url += '&hide_title=true';
      }
      if (form.cardWidth !== 500) {
        url += `&card_width=${form.cardWidth}`;
      }
      return url;
    }
  }
};
</script>

<style>
pre code {
  display: block;
  padding: 16px;
  overflow: auto;
  line-height: 1.3;
  color: #476582;
  background-color: rgba(27, 31, 35, 0.05);
  border-radius: 4px;
}

img {
  max-width: 100%;
}

.ant-input-number {
  width: 100% !important;
}
</style>