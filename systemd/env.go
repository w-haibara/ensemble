package systemd

type Env map[string]string

func NewEnv(env map[string]string) Env {
	return Env(env)
}

func (env Env) Set(key, value string) {
	env[key] = value
}

func (env Env) Unset(key string) {
	delete(env, key)
}

func (env Env) Strings() []string {
	arr := make([]string, 0, len(env))
	for k, v := range env {
		arr = append(arr, k+"="+v)
	}
	return arr
}
