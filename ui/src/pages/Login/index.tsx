import { intl } from '@/utils/intl';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { useModel, useNavigate } from '@umijs/max';
import { Button, Form, Input } from 'antd';
import React from 'react';

import { user } from '@/api';
import logoSrc from '@/assets/oceanbase_logo.svg';
import { encryptText, usePublicKey } from '@/hook/usePublicKey';
import styles from './index.less';

const Login: React.FC = () => {
  const { refresh } = useModel('@@initialState');
  const navigate = useNavigate();
  const publicKey = usePublicKey();

  const onFinish = async (values: API.User) => {
    values.password = encryptText(values.password, publicKey) as string;
    const res = await user.login(values);
    if (res.successful) {
      if (res.data.needReset) {
        navigate('/reset');
      } else {
        navigate('/overview');
      }
      refresh();
      localStorage.setItem('user', values.username);
    }
  };

  return (
    <div className={styles.loginContainer}>
      <div>
        <img className={styles.logo} src={logoSrc} alt="" />
      </div>
      <Form
        name="normal_login"
        className={styles.loginForm}
        initialValues={{ remember: true }}
        onFinish={onFinish}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: 'Please input your Username!' }]}
        >
          <Input
            style={{ width: 300 }}
            prefix={<UserOutlined />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: 'Please input your Password!' }]}
        >
          <Input
            prefix={<LockOutlined />}
            type="password"
            placeholder="Password"
          />
        </Form.Item>

        <Form.Item style={{ display: 'flex', justifyContent: 'center' }}>
          <Button style={{ width: 270 }} type="primary" htmlType="submit">
            {intl.formatMessage({
              id: 'dashboard.pages.Login.Login',
              defaultMessage: '登录',
            })}
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default Login;
