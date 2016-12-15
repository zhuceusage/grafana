///<reference path="../../../headers/common.d.ts" />

import _ from 'lodash';
import coreModule from '../../core_module';
import appEvents from '../../app_events';

export class NavBarMenuCtrl {

  /** @ngInject **/
  constructor($scope, $rootScope) {
  }
}

export function navBarMenu() {
  return {
    restrict: 'E',
    templateUrl: 'public/app/core/components/navbar/menu.html',
    controller: NavBarMenuCtrl,
    bindToController: true,
    controllerAs: 'ctrl',
    scope: {
      getMenuItems: "&",
      title: "@",
      titleUrl: "@",
      iconUrl: "@",
    },
    link: function(scope, elem) {
      // elem.click(function() {
      //   elem.toggleClass('navbar-page-btn--open');
      // });
    }
  };
}

coreModule.directive('navbarMenu', navBarMenu);
