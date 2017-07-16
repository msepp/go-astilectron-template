import { AppUiPage } from './app.po';

describe('app-ui App', () => {
  let page: AppUiPage;

  beforeEach(() => {
    page = new AppUiPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
