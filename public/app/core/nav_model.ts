

export class NavModel {

  getAlertingNav(subPage) {
    return {
      section: {
        title: 'Alerting',
        url: '/plugins',
        icon: 'icon-gf icon-gf-alert'
      },
      navItems: [
        {title: 'Alert List', active: subPage === 0, url: 'alerting/list', icon: 'fa fa-list-ul'},
        {title: 'Notifications', active: subPage === 1, url: 'alerting/notifications', icon: 'fa fa-bell-o'},
      ]
    };
  }

  getDashboardNav(dashboard, dashNav) {
    return {
      section: {
        title: dashboard.title,
        icon: 'icon-gf icon-gf-dashboard'
      },
      navItems: [
        {
          title: 'Settings',
          icon: 'fa fa-cog',
          clickHandler: () => dashNav.openEditView('settings'),
        },
        {
          title: 'Templating',
          icon: 'fa fa-code',
          clickHandler: () => dashNav.openEditView('templating'),
        },
        {
          title: 'Annotations',
          clickHandler: () => dashNav.openEditView('annotations'),
          icon: 'fa fa-bolt'
        },
        {
          title: 'View JSON',
          icon: 'fa fa-eye',
          clickHandler: () => dashNav.viewJson(),
        },
        {
          title: 'Save As',
          icon: 'fa fa-save',
          clickHandler: () => dashNav.saveDashboardAs(),
        },
      ]
    };
  }
}

var navModel = new NavModel();
export {navModel};
