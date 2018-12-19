import * as mime from 'mime-types';

/* global Cypress, cy */
// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This is will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

Cypress.Commands.add('signInAsNewUser', () => {
  // make sure we log out first before sign in
  cy.logout();

  cy.visit('/devlocal-auth/login');
  // should have both our csrf cookie tokens now
  cy.getCookie('_gorilla_csrf').should('exist');
  cy.getCookie('masked_gorilla_csrf').should('exist');
  cy.get('button[data-hook="new-user-login"]').click();
});

Cypress.Commands.add('signIntoMyMoveAsUser', userId => {
  Cypress.config('baseUrl', 'http://localhost:4000');
  cy.signInAsUser(userId);
});
Cypress.Commands.add('signIntoOffice', () => {
  Cypress.config('baseUrl', 'http://officelocal:4000');
  cy.signInAsUser('9bfa91d2-7a0c-4de0-ae02-b8cf8b4b858b');
});
Cypress.Commands.add('signIntoTSP', () => {
  Cypress.config('baseUrl', 'http://tsplocal:4000');
  cy.signInAsUser('6cd03e5b-bee8-4e97-a340-fecb8f3d5465');
});
Cypress.Commands.add('signInAsUser', userId => {
  // make sure we log out first before sign in
  cy.logout();

  cy.visit('/devlocal-auth/login');
  // should have both our csrf cookie tokens now
  cy.getCookie('_gorilla_csrf').should('exist');
  cy.getCookie('masked_gorilla_csrf').should('exist');
  cy.get('button[value="' + userId + '"]').click();
});

Cypress.Commands.add(
  'signInAsUserPostRequest',
  (
    signInAs,
    userId,
    expectedStatusCode = 200,
    expectedRespBody = null,
    sendGorillaCSRF = true,
    sendMaskedGorillaCSRF = true,
  ) => {
    // setup baseurl
    switch (signInAs.toLowerCase()) {
      case 'mymove':
        Cypress.config('baseUrl', 'http://localhost:4000');
        break;
      case 'office':
        Cypress.config('baseUrl', 'http://officelocal:4000');
        break;
      case 'tsp':
        Cypress.config('baseUrl', 'http://tsplocal:4000');
        break;
      default:
        break;
    }

    // request use to log in
    let sendRequest = maskedCSRFToken => {
      cy
        .request({
          url: '/devlocal-auth/login',
          method: 'POST',
          headers: {
            'X-CSRF-TOKEN': maskedCSRFToken,
          },
          body: {
            id: userId,
          },
          form: true,
          failOnStatusCode: false,
        })
        .then(resp => {
          cy.visit('/');
          // Default status code to check is 200
          expect(resp.status).to.eq(expectedStatusCode);
          // check response body if needed
          if (expectedRespBody) {
            expect(resp.body).to.eq(expectedRespBody);
          }
        });
    };

    // make sure we log out first before sign in
    cy.logout();
    // GET landing page to get csrf cookies
    cy.request('/');

    // Clear out cookies if we don't want to send in request
    if (!sendGorillaCSRF) {
      cy.clearCookie('_gorilla_csrf');
      // Don't include cookie in request header
      cy.getCookie('_gorilla_csrf').should('not.exist');
    } else {
      // Include cookie in request header
      cy.getCookie('_gorilla_csrf').should('exist');
    }

    if (!sendMaskedGorillaCSRF) {
      // Clear out the masked CSRF token
      cy.clearCookie('masked_gorilla_csrf');
      // Send request without masked token
      cy
        .getCookie('masked_gorilla_csrf')
        .should('not.exist')
        .then(() => {
          // null token will omit the 'X-CSRF-HEADER' from request
          sendRequest();
        });
    } else {
      // Send request with masked token
      cy
        .getCookie('masked_gorilla_csrf')
        .should('exist')
        .then(cookie => {
          sendRequest(cookie.value);
        });
    }
  },
);

Cypress.Commands.add('logout', () => {
  // The session cookie wasn't being cleared out after doing a get request even though the Set-Cookie
  // header was present. Switching to cy.visit() fixed the problem, but it's not clear why this worked.
  // Seems like others using Cypress have similar issues: https://github.com/cypress-io/cypress/issues/781
  cy.visit('/auth/logout');
});

Cypress.Commands.add('nextPage', () => {
  cy
    .get('button.next')
    .should('be.enabled')
    .click();
});

Cypress.Commands.add(
  'resetDb',
  () => {},
  /*
   * Resetting the DB in this manner is slow and should be avoided.
   * Instead of adding this to a test please create a new data record for your test in pkg/testdatagen/scenario/e2ebasic.go
   * For development you can issue `make db_e2e_reset` if you need to clean up your data.
   *
   * cy
   *   .exec('make db_e2e_reset')
   *   .its('code')
   *   .should('eq', 0),
   */
);

//from https://github.com/cypress-io/cypress/issues/669
//Cypress doesn't give the right File constructor, so we grab the window's File
Cypress.Commands.add('upload_file', (selector, fileUrl) => {
  const nameSegments = fileUrl.split('/');
  const name = nameSegments[nameSegments.length - 1];
  const rawType = mime.lookup(name);
  // mime returns false if lookup fails
  const type = rawType ? rawType : '';
  return cy.window().then(win => {
    return cy
      .fixture(fileUrl, 'base64')
      .then(Cypress.Blob.base64StringToBlob)
      .then(blob => {
        const testFile = new win.File([blob], name, { type });
        const event = {};
        event.dataTransfer = new win.DataTransfer();
        event.dataTransfer.items.add(testFile);
        return cy.get(selector).trigger('drop', event);
      });
  });
});
