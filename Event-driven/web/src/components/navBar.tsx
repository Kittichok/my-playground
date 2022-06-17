import { Layout, Menu } from 'antd';
import { Route } from 'react-router-dom';
import {
  ShoppingCartOutlined
} from '@ant-design/icons';


const { Header } = Layout;

type NavBarProps = {
  title: string;
  // component: Component;
};

const NavBar = ({ component: Component, ...rest }: any) => {
  return (
    (
      <Route
        {...rest}
        render={(props) => {
          return (
          <Layout>
            <Header>
              <div style={{ float: 'right' }}>
                <Menu
                  theme="dark"
                  mode="horizontal"
                >
                  <Menu.Item key="1">
                    <ShoppingCartOutlined />
                  </Menu.Item>
                </Menu>
              </div>
            </Header>
            <Component {...props} />
          </Layout>
          );
        }}
      />
    )
  );
}

export default NavBar;