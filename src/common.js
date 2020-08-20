class Card {
  constructor({
    width = 450,
    height = 240,
    title = "",
    body = "",
    titleHeight = 30,
    hideTitle = false,
    css = "",
    darkMode = "",
    paddingX = 25,
    paddingY = 35,
    hideBorder = false,
  }) {
    this.width = width;
    this.height = height;
    this.titleHeight = titleHeight;
    this.title = title;
    this.body = body;
    this.hideTitle = hideTitle;
    this.css = css;
    this.darkMode = darkMode;
    this.paddingX = paddingX;
    this.paddingY = paddingY;
    this.hideBorder = hideBorder;
  }

  render() {
    const cardSize = {
      width: this.width + 2*this.paddingX,
      height: this.height + 2*this.paddingY,
    };
    if(!this.hideTitle) cardSize.height += this.titleHeight;

    const bgColor = this.darkMode?"#444444":"#fffefe";
    let borderColor = "";
    if(!this.hideBorder) borderColor = this.darkMode?"#444444":"#E4E2E2";

    return `
      <svg xmlns="http://www.w3.org/2000/svg" width="${cardSize.width}" height="${cardSize.height}" viewBox="0 0 ${cardSize.width} ${cardSize.height}" fill="none">
        <style>
          .text { font: 400 11px 'Segoe UI', Ubuntu, Sans-Serif; fill: ${this.darkMode?"#fffefe":"#333333"} }
          .line { stroke:${this.darkMode?"#666666":"#dddddd"}; stroke-width:1 }
          ${this.css}
        </style>
        <rect x="0.5" y="0.5" rx="4.5" height="99%" stroke="${borderColor}" width="99%" fill="${bgColor}" stroke-opacity="1" />
        
        ${this.hideTitle ? `` : `
        <g transform="translate(${this.paddingX}, ${this.paddingY})">
          ${this.title}
        </g>`}

        <g transform="translate(${this.paddingX}, ${this.hideTitle ? this.paddingX : this.paddingX + this.paddingY})">
          ${this.body}
        </g>
      </svg>`;
  }
}

const renderError = (e, option) => {
  const css = `.t {font: 600 18px 'Microsoft Yahei UI'; fill: #e74c3c;}`
  const text = `<text class="t" dominant-baseline="text-before-edge">${e}</text>`
  return new Card({
    width: 300,
    height: 23,
    hideTitle: true,
    css,
    body: text,
    paddingY: 20,
    paddingX: 20,
    ...option,
  }).render();
};

module.exports = { Card, renderError };
