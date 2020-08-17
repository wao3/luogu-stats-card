const renderError = (e) => {
  return `
    <svg xmlns="http://www.w3.org/2000/svg" width="500" height="60" viewBox="0 0 500 60" fill="none">
      <style>.header { font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif; fill: #e74c3c;}</style>
      <rect data-testid="card-bg" x="0.5" y="0.5" rx="4.5" height="99%" stroke="#E4E2E2" width="99%" fill="#fffefe" stroke-opacity="1" />
      <g data-testid="card-title" transform="translate(25, 35)">
        <g transform="translate(0, 0)">
          <text x="0" y="0" class="header" data-testid="header">${e}</text>
        </g>
      </g>
    </svg>
      `;
};

module.exports = { renderError };
