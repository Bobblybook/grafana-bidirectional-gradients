import { e2e } from '../index';

export interface SelectOptionConfig {
  clickToOpen?: boolean;
  container: any;
  forceClickOption?: boolean;
  optionText: string | RegExp;
}

// @todo this actually returns type `Cypress.Chainable`
export const selectOption = (config: SelectOptionConfig): any => {
  const fullConfig: SelectOptionConfig = {
    clickToOpen: true,
    forceClickOption: false,
    ...config,
  };

  const { clickToOpen, container, forceClickOption, optionText } = fullConfig;

  container.within(() => {
    if (clickToOpen) {
      cy.get('[class$="-input-suffix"]', { timeout: 1000 }).then((element) => {
        expect(Cypress.dom.isAttached(element)).to.eq(true);
        cy.get('[class$="-input-suffix"]', { timeout: 1000 }).click({ force: true });
      });
    }
  });

  return e2e.components.Select.option()
    .filter((_, { textContent }) => {
      if (textContent === null) {
        return false;
      } else if (typeof optionText === 'string') {
        return textContent.includes(optionText);
      } else {
        return optionText.test(textContent);
      }
    })
    .scrollIntoView()
    .click({ force: forceClickOption });
};
